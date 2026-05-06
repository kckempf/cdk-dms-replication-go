package cdkdmsreplication


// Redshift safeguard policy for SQL Server CDC.
type SqlServerSafeguardPolicy string

const (
	SqlServerSafeguardPolicy_RELY_ON_SQL_SERVER_REPLICATION_AGENT SqlServerSafeguardPolicy = "RELY_ON_SQL_SERVER_REPLICATION_AGENT"
	SqlServerSafeguardPolicy_EXCLUSIVE_AUTOMATIC_TRUNCATION SqlServerSafeguardPolicy = "EXCLUSIVE_AUTOMATIC_TRUNCATION"
	SqlServerSafeguardPolicy_SHARED_AUTOMATIC_TRUNCATION SqlServerSafeguardPolicy = "SHARED_AUTOMATIC_TRUNCATION"
)

