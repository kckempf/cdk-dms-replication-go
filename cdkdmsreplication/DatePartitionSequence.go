package cdkdmsreplication


// Date partition sequence for S3 date-partitioned output.
type DatePartitionSequence string

const (
	DatePartitionSequence_YYYYMMDD DatePartitionSequence = "YYYYMMDD"
	DatePartitionSequence_YYYYMMDDHH DatePartitionSequence = "YYYYMMDDHH"
	DatePartitionSequence_YYYYMM DatePartitionSequence = "YYYYMM"
	DatePartitionSequence_MMYYYYDD DatePartitionSequence = "MMYYYYDD"
	DatePartitionSequence_DDMMYYYY DatePartitionSequence = "DDMMYYYY"
)

