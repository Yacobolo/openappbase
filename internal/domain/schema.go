// internal/domain/schema.go
package domain

type Column struct {
	Name          string
	DataType      string
	IsNullable    bool
	ColumnDefault string
}

type Table struct {
	Name               string
	Owner              string
	RowSecurityEnabled bool
	EstimatedRows      int
}

type Schema struct {
	Name   string
	Tables []Table
}

// TableData represents the result of fetching data from a table
type TableData struct {
	Columns    []Column
	Rows       [][]string // Each row is a slice of string values
	Pagination PaginationInfo
}

// PaginationInfo contains metadata about pagination state
type PaginationInfo struct {
	CurrentPage int
	TotalPages  int
	PageSize    int
	TotalRows   int64
	HasPrevious bool
	HasNext     bool
}
