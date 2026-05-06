package cdkdmsreplication


// Target DB type for MySQL target endpoints.
type MySqlTargetDbType string

const (
	MySqlTargetDbType_SPECIFIC_DATABASE MySqlTargetDbType = "SPECIFIC_DATABASE"
	MySqlTargetDbType_MULTIPLE_DATABASES MySqlTargetDbType = "MULTIPLE_DATABASES"
)

