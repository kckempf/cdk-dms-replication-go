package cdkdmsreplication


// Controls CDC behaviour.
type CdcSettings struct {
	// Whether DMS applies changes immediately (false) or batches them (true).
	// Default: false.
	//
	BatchApplyEnabled *bool `field:"optional" json:"batchApplyEnabled" yaml:"batchApplyEnabled"`
	// Memory limit (in MB) for the apply task.
	// Default: 500.
	//
	BatchApplyMemoryLimit *float64 `field:"optional" json:"batchApplyMemoryLimit" yaml:"batchApplyMemoryLimit"`
	// Maximum number of seconds DMS waits before applying batched changes.
	// Default: 30.
	//
	BatchApplyTimeoutMax *float64 `field:"optional" json:"batchApplyTimeoutMax" yaml:"batchApplyTimeoutMax"`
	// Number of seconds DMS waits before applying batched changes.
	//
	// Only relevant when `batchApplyEnabled` is true.
	// Default: 1.
	//
	BatchApplyTimeoutMin *float64 `field:"optional" json:"batchApplyTimeoutMin" yaml:"batchApplyTimeoutMin"`
	// Offset in seconds from current time to use as the CDC start position.
	//
	// Ignored when `cdcStartPosition` is set on the task.
	EnableTestMode *bool `field:"optional" json:"enableTestMode" yaml:"enableTestMode"`
	// Whether to preserve transactions during CDC.
	// Default: true.
	//
	FailOnNoTablesCaptured *bool `field:"optional" json:"failOnNoTablesCaptured" yaml:"failOnNoTablesCaptured"`
}

