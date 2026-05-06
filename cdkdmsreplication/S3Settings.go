package cdkdmsreplication


// Settings for Amazon S3 endpoints (source or target).
type S3Settings struct {
	// S3 bucket name.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// IAM role ARN that DMS uses to access the S3 bucket.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Whether DMS adds a column name field to CSV output.
	AddColumnName *bool `field:"optional" json:"addColumnName" yaml:"addColumnName"`
	// Folder path prefix within the bucket.
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// Include CDC inserts and updates in the target.
	CdcInsertsAndUpdates *bool `field:"optional" json:"cdcInsertsAndUpdates" yaml:"cdcInsertsAndUpdates"`
	// Include only inserts (not updates or deletes) in the target.
	CdcInsertsOnly *bool `field:"optional" json:"cdcInsertsOnly" yaml:"cdcInsertsOnly"`
	// Maximum interval in seconds between CDC mini-batches (1-360).
	CdcMaxBatchInterval *float64 `field:"optional" json:"cdcMaxBatchInterval" yaml:"cdcMaxBatchInterval"`
	// Minimum file size (in KB) to trigger a CDC file write.
	CdcMinFileSize *float64 `field:"optional" json:"cdcMinFileSize" yaml:"cdcMinFileSize"`
	// CDC path in the source S3 bucket.
	CdcPath *string `field:"optional" json:"cdcPath" yaml:"cdcPath"`
	// Column delimiter character for CSV output.
	//
	// Default: comma.
	CsvDelimiter *string `field:"optional" json:"csvDelimiter" yaml:"csvDelimiter"`
	// String used for null values when no-sup-value applies.
	CsvNoSupValue *string `field:"optional" json:"csvNoSupValue" yaml:"csvNoSupValue"`
	// String used for null values in CSV output.
	CsvNullValue *string `field:"optional" json:"csvNullValue" yaml:"csvNullValue"`
	// Row delimiter for CSV output.
	//
	// Default: newline.
	CsvRowDelimiter *string `field:"optional" json:"csvRowDelimiter" yaml:"csvRowDelimiter"`
	// Output data format: CSV or Parquet.
	DataFormat S3DataFormat `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Date partition delimiter.
	DatePartitionDelimiter DatePartitionDelimiter `field:"optional" json:"datePartitionDelimiter" yaml:"datePartitionDelimiter"`
	// Whether to partition output files by date.
	DatePartitionEnabled *bool `field:"optional" json:"datePartitionEnabled" yaml:"datePartitionEnabled"`
	// Date partition sequence.
	DatePartitionSequence DatePartitionSequence `field:"optional" json:"datePartitionSequence" yaml:"datePartitionSequence"`
	// Dictionary page size (in bytes) for Parquet.
	DictPageSizeLimit *float64 `field:"optional" json:"dictPageSizeLimit" yaml:"dictPageSizeLimit"`
	// Whether to enable statistics for Parquet row groups.
	EnableStatistics *bool `field:"optional" json:"enableStatistics" yaml:"enableStatistics"`
	// Encoding type for Parquet.
	EncodingType EncodingType `field:"optional" json:"encodingType" yaml:"encodingType"`
	// Encryption mode for data at rest in S3.
	EncryptionMode EncryptionMode `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// JSON structure defining external tables (for S3 source).
	ExternalTableDefinition *string `field:"optional" json:"externalTableDefinition" yaml:"externalTableDefinition"`
	// Number of header rows to ignore in the S3 source.
	IgnoreHeaderRows *float64 `field:"optional" json:"ignoreHeaderRows" yaml:"ignoreHeaderRows"`
	// Whether to include the operation column for full-load rows.
	IncludeOpForFullLoad *bool `field:"optional" json:"includeOpForFullLoad" yaml:"includeOpForFullLoad"`
	// Maximum file size (in KB) for data files.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Whether timestamps in Parquet use milliseconds instead of microseconds.
	ParquetTimestampInMillisecond *bool `field:"optional" json:"parquetTimestampInMillisecond" yaml:"parquetTimestampInMillisecond"`
	// Parquet format version.
	ParquetVersion ParquetVersion `field:"optional" json:"parquetVersion" yaml:"parquetVersion"`
	// Whether DMS preserves transaction boundaries (for CDC).
	PreserveTransactions *bool `field:"optional" json:"preserveTransactions" yaml:"preserveTransactions"`
	// Whether to use RFC 4180 compliant CSV format.
	Rfc4180 *bool `field:"optional" json:"rfc4180" yaml:"rfc4180"`
	// Number of rows per Parquet row group.
	RowGroupLength *float64 `field:"optional" json:"rowGroupLength" yaml:"rowGroupLength"`
	// KMS key ARN for SSE-KMS encryption.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// Column name for operation timestamps.
	TimestampColumnName *string `field:"optional" json:"timestampColumnName" yaml:"timestampColumnName"`
	// Whether to use CSV no-sup-value.
	UseCsvNoSupValue *bool `field:"optional" json:"useCsvNoSupValue" yaml:"useCsvNoSupValue"`
	// Whether to use the task start time instead of transaction start time for full load.
	UseTaskStartTimeForFullLoadTimestamp *bool `field:"optional" json:"useTaskStartTimeForFullLoadTimestamp" yaml:"useTaskStartTimeForFullLoadTimestamp"`
}

