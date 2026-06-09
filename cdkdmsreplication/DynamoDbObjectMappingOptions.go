package cdkdmsreplication


// Options for `TableMappings.mapToDynamoDb`.
type DynamoDbObjectMappingOptions struct {
	// Mapping for the DynamoDB partition (hash) key.
	//
	// Required.
	PartitionKey *DynamoDbKeyMapping `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Name of the target DynamoDB table.
	TargetTableName *string `field:"required" json:"targetTableName" yaml:"targetTableName"`
	// Additional column-to-attribute mappings.
	//
	// Use this to rename non-key
	// columns or change their DynamoDB sub-type. Columns not listed here pass
	// through with the source column name.
	AttributeMappings *[]*DynamoDbAttributeMapping `field:"optional" json:"attributeMappings" yaml:"attributeMappings"`
	// Source columns to exclude from the migrated item.
	ExcludeColumns *[]*string `field:"optional" json:"excludeColumns" yaml:"excludeColumns"`
	// Mapping for the DynamoDB sort (range) key.
	//
	// Optional.
	SortKey *DynamoDbKeyMapping `field:"optional" json:"sortKey" yaml:"sortKey"`
}

