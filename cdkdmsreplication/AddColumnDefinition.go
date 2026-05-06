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
	// Constant value to populate the column with.
	//
	// Exactly one of `columnValue` or `expression` must be set.
	ColumnValue *string `field:"optional" json:"columnValue" yaml:"columnValue"`
	// Expression (e.g. `$timestamp`) to populate the column. Exactly one of `columnValue` or `expression` must be set.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

