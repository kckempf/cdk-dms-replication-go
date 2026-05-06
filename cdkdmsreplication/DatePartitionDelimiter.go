package cdkdmsreplication


// Date partition delimiter for S3 date-partitioned output.
type DatePartitionDelimiter string

const (
	DatePartitionDelimiter_SLASH DatePartitionDelimiter = "SLASH"
	DatePartitionDelimiter_UNDERSCORE DatePartitionDelimiter = "UNDERSCORE"
	DatePartitionDelimiter_DASH DatePartitionDelimiter = "DASH"
	DatePartitionDelimiter_NONE DatePartitionDelimiter = "NONE"
)

