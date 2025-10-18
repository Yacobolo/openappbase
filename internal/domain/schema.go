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
