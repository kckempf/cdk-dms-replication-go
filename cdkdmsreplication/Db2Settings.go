package cdkdmsreplication


// Settings for IBM Db2 LUW endpoints.
type Db2Settings struct {
	// Current LSN as of which DMS should start reading.
	CurrentLsn *string `field:"optional" json:"currentLsn" yaml:"currentLsn"`
	// (Target) Whether DMS keeps the .csv files it used to load data, for troubleshooting.
	KeepCsvFiles *bool `field:"optional" json:"keepCsvFiles" yaml:"keepCsvFiles"`
	// (Target) Time in milliseconds DMS waits before timing out a data load.
	LoadTimeout *float64 `field:"optional" json:"loadTimeout" yaml:"loadTimeout"`
	// (Target) Maximum size in KB of the .csv files DMS uses to load data.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Maximum number of bytes per read operation.
	MaxKBytesPerRead *float64 `field:"optional" json:"maxKBytesPerRead" yaml:"maxKBytesPerRead"`
	// ARN of IAM role for Secrets Manager.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Enables ongoing replication (CDC) for Db2.
	SetDataCaptureChanges *bool `field:"optional" json:"setDataCaptureChanges" yaml:"setDataCaptureChanges"`
	// (Target) Size in KB of the in-memory buffer DMS writes to before flushing to a .csv file.
	WriteBufferSize *float64 `field:"optional" json:"writeBufferSize" yaml:"writeBufferSize"`
}

