package cdkdmsreplication


// Settings for Amazon Kinesis Data Streams target endpoints.
type KinesisSettings struct {
	// ARN of the IAM role that provides DMS access to Kinesis.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// ARN of the Kinesis stream.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// Whether to include control details in messages.
	IncludeControlDetails *bool `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Whether to include null and empty columns in messages.
	IncludeNullAndEmpty *bool `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Whether to include the partition value in messages.
	IncludePartitionValue *bool `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Whether to include table ALTER operations in messages.
	IncludeTableAlterOperations *bool `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Whether to include transaction details in messages.
	IncludeTransactionDetails *bool `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// Message format.
	MessageFormat MessageFormat `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Whether to omit the hex prefix from binary values.
	NoHexPrefix *bool `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Whether to include the schema name in the partition key.
	PartitionIncludeSchemaTable *bool `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
}

