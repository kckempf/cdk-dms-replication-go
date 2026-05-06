package cdkdmsreplication


// Settings for IBM Db2 LUW endpoints.
type Db2Settings struct {
	// Current LSN as of which DMS should start reading.
	CurrentLsn *string `field:"optional" json:"currentLsn" yaml:"currentLsn"`
	// Maximum number of bytes per read operation.
	MaxKBytesPerRead *float64 `field:"optional" json:"maxKBytesPerRead" yaml:"maxKBytesPerRead"`
	// ARN of IAM role for Secrets Manager.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Enables ongoing replication (CDC) for Db2.
	SetDataCaptureChanges *bool `field:"optional" json:"setDataCaptureChanges" yaml:"setDataCaptureChanges"`
}

