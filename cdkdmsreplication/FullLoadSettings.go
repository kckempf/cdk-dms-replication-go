package cdkdmsreplication


// Controls how the task handles full-load data.
type FullLoadSettings struct {
	// Whether DMS commits the full load in a single transaction.
	CommitRate *float64 `field:"optional" json:"commitRate" yaml:"commitRate"`
	// Maximum number of tables to load in parallel.
	// Default: 8.
	//
	MaxFullLoadSubTasks *float64 `field:"optional" json:"maxFullLoadSubTasks" yaml:"maxFullLoadSubTasks"`
	// Whether to stop the task after the full load completes (only relevant for full-load tasks without CDC).
	// Default: false.
	//
	StopTaskCachedChangesApplied *bool `field:"optional" json:"stopTaskCachedChangesApplied" yaml:"stopTaskCachedChangesApplied"`
	// Whether to create target tables if they don't exist.
	// Default: true.
	//
	TargetTablePrepMode *string `field:"optional" json:"targetTablePrepMode" yaml:"targetTablePrepMode"`
	// Number of rows to load before a commit.
	// Default: 10000.
	//
	TransactionConsistencyTimeout *float64 `field:"optional" json:"transactionConsistencyTimeout" yaml:"transactionConsistencyTimeout"`
}

