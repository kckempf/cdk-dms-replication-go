package cdkdmsreplication


// Data format for S3 endpoint output.
type S3DataFormat string

const (
	S3DataFormat_CSV S3DataFormat = "CSV"
	S3DataFormat_PARQUET S3DataFormat = "PARQUET"
)

