// internal/features/explorer/services/explorer_service.go
package services

import (
	"context"
	"northstar/internal/domain"
	"northstar/internal/store"

	"github.com/jackc/pgx/v5/pgtype"
)

type ExplorerService struct {
	q *store.Queries
}

func NewExplorerService(q *store.Queries) *ExplorerService {
	return &ExplorerService{
		q: q,
	}
}

func (s *ExplorerService) GetSchemaOverview(ctx context.Context) ([]domain.Schema, error) {
	rows, err := s.q.GetFullSchemaDetails(ctx)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}

	var domainSchemas []domain.Schema

	// Start with the first schema from our sorted list.
	currentSchema := domain.Schema{
		Name:   rows[0].TableSchema,
		Tables: []domain.Table{},
	}

	for _, row := range rows {
		// Since the list is sorted by schema, a name change
		// means we've finished the previous schema.
		if row.TableSchema != currentSchema.Name {
			// Add the completed schema to our final slice.
			domainSchemas = append(domainSchemas, currentSchema)

			// And start the next one.
			currentSchema = domain.Schema{
				Name:   row.TableSchema,
				Tables: []domain.Table{},
			}
		}

		// Add the current table to the schema we're building.
		currentSchema.Tables = append(currentSchema.Tables, domain.Table{
			Name: row.TableName,
		})
	}

	// Don't forget to add the very last schema being built.
	domainSchemas = append(domainSchemas, currentSchema)

	return domainSchemas, nil
}

func (s *ExplorerService) GetTableColumns(ctx context.Context, schemaName, tableName string) ([]domain.Column, error) {
	params := store.GetTableColumnsParams{
		TableSchema: pgtype.Text{String: schemaName, Valid: true},
		TableName:   pgtype.Text{String: tableName, Valid: true},
	}
	rows, err := s.q.GetTableColumns(ctx, params)
	if err != nil {
		return nil, err
	}

	var columns []domain.Column
	for _, row := range rows {
		columns = append(columns, domain.Column{
			Name:          row.ColumnName.(string),
			DataType:      row.DataType.(string),
			IsNullable:    row.IsNullable,
			ColumnDefault: row.ColumnDefault.(string),
		})
	}

	return columns, nil
}
