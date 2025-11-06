// internal/features/editor/services/editor_service.go
package services

import (
	"context"
	"fmt"
	"math"
	"northstar/internal/domain"
	"northstar/internal/store"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

const PageSize = 50

type EditorService struct {
	q    *store.Queries
	pool *pgxpool.Pool
}

func NewEditorService(q *store.Queries, pool *pgxpool.Pool) *EditorService {
	return &EditorService{
		q:    q,
		pool: pool,
	}
}

// ParseTableName splits "schema.table" format into schema and table names
func (s *EditorService) ParseTableName(tableName string) (schema, table string, err error) {
	parts := strings.Split(tableName, ".")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid table name format: expected 'schema.table', got '%s'", tableName)
	}
	return parts[0], parts[1], nil
}

// ValidateTable checks if a table exists in the database
func (s *EditorService) ValidateTable(ctx context.Context, schema, table string) (bool, error) {
	params := store.ValidateTableExistsParams{
		TableSchema: pgtype.Text{String: schema, Valid: true},
		TableName:   pgtype.Text{String: table, Valid: true},
	}
	return s.q.ValidateTableExists(ctx, params)
}

// GetTableRowCount returns the total number of rows in a table
func (s *EditorService) GetTableRowCount(ctx context.Context, schema, table string) (int64, error) {
	// Use format with %I to safely escape identifiers
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", pgx.Identifier{schema}.Sanitize(), pgx.Identifier{table}.Sanitize())

	var count int64
	err := s.pool.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count rows in %s.%s: %w", schema, table, err)
	}

	return count, nil
}

// GetTableData fetches paginated data from a table
func (s *EditorService) GetTableData(ctx context.Context, schema, table string, page int) (*domain.TableData, error) {
	// Validate table exists first
	exists, err := s.ValidateTable(ctx, schema, table)
	if err != nil {
		return nil, fmt.Errorf("failed to validate table: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("table %s.%s does not exist", schema, table)
	}

	// Get column information
	columns, err := s.getTableColumns(ctx, schema, table)
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	// Get total row count
	totalRows, err := s.GetTableRowCount(ctx, schema, table)
	if err != nil {
		return nil, fmt.Errorf("failed to get row count: %w", err)
	}

	// Calculate pagination
	totalPages := int(math.Ceil(float64(totalRows) / float64(PageSize)))
	if page < 1 {
		page = 1
	}
	if page > totalPages && totalPages > 0 {
		page = totalPages
	}

	offset := (page - 1) * PageSize

	// Fetch data rows
	rows, err := s.getTableRows(ctx, schema, table, PageSize, offset, columns)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows: %w", err)
	}

	// Build pagination info
	pagination := domain.PaginationInfo{
		CurrentPage: page,
		TotalPages:  totalPages,
		PageSize:    PageSize,
		TotalRows:   totalRows,
		HasPrevious: page > 1,
		HasNext:     page < totalPages,
	}

	return &domain.TableData{
		Columns:    columns,
		Rows:       rows,
		Pagination: pagination,
	}, nil
}

// getTableColumns retrieves column metadata using existing query
func (s *EditorService) getTableColumns(ctx context.Context, schema, table string) ([]domain.Column, error) {
	params := store.GetTableColumnsParams{
		TableSchema: pgtype.Text{String: schema, Valid: true},
		TableName:   pgtype.Text{String: table, Valid: true},
	}

	rows, err := s.q.GetTableColumns(ctx, params)
	if err != nil {
		return nil, err
	}

	columns := make([]domain.Column, len(rows))
	for i, row := range rows {
		columnDefault := ""
		if row.ColumnDefault.Valid {
			columnDefault = row.ColumnDefault.String
		}
		columns[i] = domain.Column{
			Name:          row.ColumnName,
			DataType:      row.DataType,
			IsNullable:    row.IsNullable,
			ColumnDefault: columnDefault,
		}
	}

	return columns, nil
}

// getTableRows fetches actual data rows from the table
func (s *EditorService) getTableRows(ctx context.Context, schema, table string, limit, offset int, columns []domain.Column) ([][]string, error) {
	// Build column list for SELECT
	columnNames := make([]string, len(columns))
	for i, col := range columns {
		columnNames[i] = pgx.Identifier{col.Name}.Sanitize()
	}
	columnsStr := strings.Join(columnNames, ", ")

	// Build query with proper identifier escaping
	query := fmt.Sprintf(
		"SELECT %s FROM %s.%s ORDER BY 1 LIMIT $1 OFFSET $2",
		columnsStr,
		pgx.Identifier{schema}.Sanitize(),
		pgx.Identifier{table}.Sanitize(),
	)

	rows, err := s.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query table data: %w", err)
	}
	defer rows.Close()

	var result [][]string
	for rows.Next() {
		// Get raw values
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert all values to strings
		rowData := make([]string, len(values))
		for i, val := range values {
			if val == nil {
				rowData[i] = ""
			} else {
				rowData[i] = fmt.Sprintf("%v", val)
			}
		}
		result = append(result, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return result, nil
}
