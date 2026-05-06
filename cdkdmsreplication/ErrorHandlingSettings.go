package cdkdmsreplication


// Controls error handling behaviour.
type ErrorHandlingSettings struct {
	// Maximum number of data errors before stopping.
	// Default: 0 (unlimited).
	//
	DataErrorEscalationCount *float64 `field:"optional" json:"dataErrorEscalationCount" yaml:"dataErrorEscalationCount"`
	// Action to take after `dataErrorEscalationCount` errors are hit.
	// Default: ErrorAction.STOP_TASK
	//
	DataErrorEscalationPolicy ErrorAction `field:"optional" json:"dataErrorEscalationPolicy" yaml:"dataErrorEscalationPolicy"`
	// Action to take when DMS encounters a data error (e.g. duplicate key).
	// Default: ErrorAction.STOP_TASK
	//
	DataErrorPolicy ErrorAction `field:"optional" json:"dataErrorPolicy" yaml:"dataErrorPolicy"`
	// Whether to recover from recoverable errors.
	// Default: true.
	//
	RecoverableErrorCount *float64 `field:"optional" json:"recoverableErrorCount" yaml:"recoverableErrorCount"`
	// Interval in seconds between recovery attempts.
	// Default: 5.
	//
	RecoverableErrorInterval *float64 `field:"optional" json:"recoverableErrorInterval" yaml:"recoverableErrorInterval"`
	// Throttle factor applied to recovery intervals.
	// Default: 2.
	//
	RecoverableErrorThrottling *bool `field:"optional" json:"recoverableErrorThrottling" yaml:"recoverableErrorThrottling"`
	// Maximum interval (seconds) between recovery attempts.
	// Default: 360.
	//
	RecoverableErrorThrottlingMax *float64 `field:"optional" json:"recoverableErrorThrottlingMax" yaml:"recoverableErrorThrottlingMax"`
	// Maximum number of table errors before escalation.
	// Default: 0.
	//
	TableErrorEscalationCount *float64 `field:"optional" json:"tableErrorEscalationCount" yaml:"tableErrorEscalationCount"`
	// Action after `tableErrorEscalationCount` table errors.
	// Default: ErrorAction.STOP_TASK
	//
	TableErrorEscalationPolicy ErrorAction `field:"optional" json:"tableErrorEscalationPolicy" yaml:"tableErrorEscalationPolicy"`
	// Action when DMS encounters a table error.
	// Default: ErrorAction.STOP_TASK
	//
	TableErrorPolicy ErrorAction `field:"optional" json:"tableErrorPolicy" yaml:"tableErrorPolicy"`
}

