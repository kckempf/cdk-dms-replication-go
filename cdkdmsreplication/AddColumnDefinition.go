package cdkdmsreplication


// A column added via an ADD_COLUMN transformation.
type AddColumnDefinition struct {
	// Name of the new column.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Data type of the new column.
	ColumnType ColumnDataType `field:"required" json:"columnType" yaml:"columnType"`
	// Character length for string columns.
	ColumnLength *float64 `field:"optional" json:"columnLength" yaml:"columnLength"`
	// Precision for numeric columns.
	ColumnPrecision *float64 `field:"optional" json:"columnPrecision" yaml:"columnPrecision"`
	// Scale for numeric columns.
	ColumnScale *float64 `field:"optional" json:"columnScale" yaml:"columnScale"`
	// Constant string value to populate the column with.
	//
	// Emitted as a
	// single-quoted DMS expression literal — use only for string column types.
	// For numeric or datetime types, use `expression` with an unquoted literal
	// (e.g. `expression: '42'`).
	// Exactly one of `columnValue` or `expression` must be set.
	ColumnValue *string `field:"optional" json:"columnValue" yaml:"columnValue"`
	// DMS expression to populate the column (e.g. `$timestamp`, `'ENTITY#' || $id`, or an unquoted numeric literal like `42`). Exactly one of `columnValue` or `expression` must be set.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

