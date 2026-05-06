package cdkdmsreplication


// Object locator identifying the schema, table, and optional column a rule targets.
type RuleObjectLocatorValue struct {
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

