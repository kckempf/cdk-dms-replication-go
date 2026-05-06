package cdkdmsreplication


// Logging verbosity for DMS task log components.
type LoggingLevel string

const (
	LoggingLevel_LOGGER_SEVERITY_DEFAULT LoggingLevel = "LOGGER_SEVERITY_DEFAULT"
	LoggingLevel_LOGGER_SEVERITY_DEBUG LoggingLevel = "LOGGER_SEVERITY_DEBUG"
	LoggingLevel_LOGGER_SEVERITY_DETAILED_DEBUG LoggingLevel = "LOGGER_SEVERITY_DETAILED_DEBUG"
	LoggingLevel_LOGGER_SEVERITY_ERROR LoggingLevel = "LOGGER_SEVERITY_ERROR"
	LoggingLevel_LOGGER_SEVERITY_WARNING LoggingLevel = "LOGGER_SEVERITY_WARNING"
)

