package cdkdmsreplication


// Settings for Amazon Neptune target endpoints.
type NeptuneSettings struct {
	// Folder within the S3 bucket for Neptune staging data.
	S3BucketFolder *string `field:"required" json:"s3BucketFolder" yaml:"s3BucketFolder"`
	// S3 bucket where DMS stages the migration data.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// ARN of the IAM role that provides DMS access to S3 and Neptune.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Number of seconds to retry on errors before failing the task.
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// Whether IAM auth is enabled on the Neptune cluster.
	IamAuthEnabled *bool `field:"optional" json:"iamAuthEnabled" yaml:"iamAuthEnabled"`
	// Maximum number of files per request to the Neptune bulk-load API.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Maximum number of retries on endpoint exceptions.
	MaxRetryCount *float64 `field:"optional" json:"maxRetryCount" yaml:"maxRetryCount"`
}

