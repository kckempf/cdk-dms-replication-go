package cdkdmsreplication


// Settings for Amazon OpenSearch Service target endpoints.
type OpenSearchSettings struct {
	// Endpoint URL of the OpenSearch cluster (e.g. https://my-domain.us-east-1.es.amazonaws.com).
	EndpointUri *string `field:"required" json:"endpointUri" yaml:"endpointUri"`
	// ARN of the IAM role that provides DMS access to OpenSearch.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Number of seconds to retry on errors before failing.
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// Maximum percentage of records that may fail before the task is stopped.
	FullLoadErrorPercentage *float64 `field:"optional" json:"fullLoadErrorPercentage" yaml:"fullLoadErrorPercentage"`
}

