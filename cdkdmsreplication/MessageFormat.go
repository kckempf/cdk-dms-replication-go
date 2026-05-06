package cdkdmsreplication


// Message format emitted to Kinesis or Kafka.
type MessageFormat string

const (
	MessageFormat_JSON MessageFormat = "JSON"
	MessageFormat_JSON_UNFORMATTED MessageFormat = "JSON_UNFORMATTED"
)

