package cdkdmsreplication


// Action to take on a recoverable error.
type ErrorAction string

const (
	// Ignore the error and continue.
	ErrorAction_IGNORE ErrorAction = "IGNORE"
	// Ignore the row, continue with the next.
	ErrorAction_IGNORE_RECORD ErrorAction = "IGNORE_RECORD"
	// Stop the task.
	ErrorAction_STOP_TASK ErrorAction = "STOP_TASK"
)

