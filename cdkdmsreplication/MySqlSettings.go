package cdkdmsreplication


// Settings for MySQL and MariaDB endpoints.
type MySqlSettings struct {
	// SQL to run after connecting to the endpoint.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Remove metadata from source that differs from target during full load.
	CleanSourceMetadataOnMismatch *bool `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Interval in seconds between polls for source events during CDC.
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Maximum file size (in KB) for CSV files created during full load.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Number of parallel threads to use for a full load.
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// ARN of the IAM role that provides access to the Secrets Manager secret.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the secret in AWS Secrets Manager containing the endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Specifies the time zone for source MySQL.
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// For a MySQL target, specifies how tables are created.
	TargetDbType MySqlTargetDbType `field:"optional" json:"targetDbType" yaml:"targetDbType"`
}

