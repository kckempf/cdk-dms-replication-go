package cdkdmsreplication


// The type of database migration to perform.
type MigrationType string

const (
	// Migrate all existing data from source to target (one-time).
	MigrationType_FULL_LOAD MigrationType = "FULL_LOAD"
	// Replicate ongoing changes from source to target using CDC.
	MigrationType_CDC MigrationType = "CDC"
	// Perform a full load first, then replicate ongoing changes.
	MigrationType_FULL_LOAD_AND_CDC MigrationType = "FULL_LOAD_AND_CDC"
)

