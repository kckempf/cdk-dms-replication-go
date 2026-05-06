package cdkdmsreplication


// Logging configuration for a single DMS log component.
type LogComponentSettings struct {
	// Logging level for this component.
	LoggingLevel LoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
}

