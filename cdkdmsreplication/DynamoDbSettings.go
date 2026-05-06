package cdkdmsreplication


// Settings for Amazon DynamoDB target endpoints.
type DynamoDbSettings struct {
	// ARN of the IAM role that provides DMS access to DynamoDB.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

