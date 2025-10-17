// store/store.go
package store

import (
	"fmt"
	"log"
)

type TableInfo struct {
	SchemaName  string      `json:"schemaname"`
	TableName   string      `json:"tablename"`
	TableOwner  string      `json:"tableowner"`
	Tablespace  interface{} `json:"tablespace"`
	HasIndexes  bool        `json:"hasindexes"`
	HasRules    bool        `json:"hasrules"`
	HasTriggers bool        `json:"hastriggers"`
	RowSecurity bool        `json:"rowsecurity"`
}

func GetAllTablesInfo() ([]TableInfo, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	// Query to get detailed information for all tables.
	query := `
		SELECT 
			schemaname, 
			tablename, 
			tableowner, 
			tablespace, 
			hasindexes, 
			hasrules, 
			hastriggers, 
			rowsecurity 
		FROM pg_catalog.pg_tables
		ORDER BY schemaname, tablename;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query for tables info: %w", err)
	}
	defer rows.Close()

	var tables []TableInfo
	for rows.Next() {
		var t TableInfo
		if err := rows.Scan(
			&t.SchemaName,
			&t.TableName,
			&t.TableOwner,
			&t.Tablespace,
			&t.HasIndexes,
			&t.HasRules,
			&t.HasTriggers,
			&t.RowSecurity,
		); err != nil {
			return nil, fmt.Errorf("failed to scan table info: %w", err)
		}
		tables = append(tables, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration for tables info: %w", err)
	}

	return tables, nil
}

func GetTableData(tableName string) ([]map[string]interface{}, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	// Build and execute the query
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}
	columnCount := len(columns)

	// Create a slice of maps to hold the results
	var results []map[string]interface{}

	values := make([]interface{}, columnCount)
	valuePtrs := make([]interface{}, columnCount)

	// Iterate over rows
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			continue
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]

			// Convert []byte to string for better handling (e.g., JSON marshaling)
			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}
		results = append(results, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return results, nil
}
