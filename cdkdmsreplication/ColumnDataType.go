package cdkdmsreplication


// Data type for added columns.
type ColumnDataType string

const (
	ColumnDataType_STRING ColumnDataType = "STRING"
	ColumnDataType_INT4 ColumnDataType = "INT4"
	ColumnDataType_INT8 ColumnDataType = "INT8"
	ColumnDataType_FLOAT4 ColumnDataType = "FLOAT4"
	ColumnDataType_FLOAT8 ColumnDataType = "FLOAT8"
	ColumnDataType_NUMERIC ColumnDataType = "NUMERIC"
	ColumnDataType_DATETIME ColumnDataType = "DATETIME"
	ColumnDataType_BYTES ColumnDataType = "BYTES"
	ColumnDataType_BLOB ColumnDataType = "BLOB"
)

