package cdkdmsreplication


// Settings for Amazon Redshift target endpoints.
type RedshiftSettings struct {
	// Whether to accept dates in a specific format.
	AcceptAnyDate *bool `field:"optional" json:"acceptAnyDate" yaml:"acceptAnyDate"`
	// SQL to run after connecting.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// S3 intermediate bucket folder path.
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// S3 bucket name used for the intermediate storage.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Whether to enable case-sensitive column names.
	CaseSensitiveNames *bool `field:"optional" json:"caseSensitiveNames" yaml:"caseSensitiveNames"`
	// Whether to enable automatic compression.
	CompUpdate *bool `field:"optional" json:"compUpdate" yaml:"compUpdate"`
	// Timeout in seconds for database connections.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// Date format for the DATEFORMAT option.
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// Whether to load empty strings as NULL.
	EmptyAsNull *bool `field:"optional" json:"emptyAsNull" yaml:"emptyAsNull"`
	// Encryption mode for data at rest.
	EncryptionMode EncryptionMode `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Whether to allow explicit ID values in the COPY command.
	ExplicitIds *bool `field:"optional" json:"explicitIds" yaml:"explicitIds"`
	// Number of upload streams for parallel loading.
	FileTransferUploadStreams *float64 `field:"optional" json:"fileTransferUploadStreams" yaml:"fileTransferUploadStreams"`
	// Timeout (in seconds) for loading data.
	LoadTimeout *float64 `field:"optional" json:"loadTimeout" yaml:"loadTimeout"`
	// Maximum file size (in KB) for each intermediate file.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Whether to remove quoted data.
	RemoveQuotes *bool `field:"optional" json:"removeQuotes" yaml:"removeQuotes"`
	// Character to replace a specific character with.
	ReplaceChars *string `field:"optional" json:"replaceChars" yaml:"replaceChars"`
	// Character to use to replace invalid UTF-8 characters.
	ReplaceInvalidChars *string `field:"optional" json:"replaceInvalidChars" yaml:"replaceInvalidChars"`
	// ARN of IAM role for accessing Secrets Manager.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// KMS key ARN for SSE-KMS.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// ARN of the IAM role that provides DMS access to the S3 staging bucket.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Date format for the TIMEFORMAT option.
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// Whether to remove trailing blanks from VARCHAR columns.
	TrimBlanks *bool `field:"optional" json:"trimBlanks" yaml:"trimBlanks"`
	// Whether to truncate VARCHAR columns to the maximum length.
	TruncateColumns *bool `field:"optional" json:"truncateColumns" yaml:"truncateColumns"`
	// Size of the write buffer (in KB).
	WriteBufferSize *float64 `field:"optional" json:"writeBufferSize" yaml:"writeBufferSize"`
}

