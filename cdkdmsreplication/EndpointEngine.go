package cdkdmsreplication


// Database engine for a DMS endpoint.
type EndpointEngine string

const (
	EndpointEngine_MYSQL EndpointEngine = "MYSQL"
	EndpointEngine_AURORA_MYSQL EndpointEngine = "AURORA_MYSQL"
	EndpointEngine_POSTGRES EndpointEngine = "POSTGRES"
	EndpointEngine_AURORA_POSTGRESQL EndpointEngine = "AURORA_POSTGRESQL"
	EndpointEngine_ORACLE EndpointEngine = "ORACLE"
	EndpointEngine_SQLSERVER EndpointEngine = "SQLSERVER"
	EndpointEngine_MARIADB EndpointEngine = "MARIADB"
	EndpointEngine_SAP_ASE EndpointEngine = "SAP_ASE"
	EndpointEngine_IBM_DB2 EndpointEngine = "IBM_DB2"
	EndpointEngine_IBM_DB2_ZOS EndpointEngine = "IBM_DB2_ZOS"
	EndpointEngine_MONGODB EndpointEngine = "MONGODB"
	EndpointEngine_DOCDB EndpointEngine = "DOCDB"
	EndpointEngine_S3 EndpointEngine = "S3"
	EndpointEngine_DYNAMODB EndpointEngine = "DYNAMODB"
	EndpointEngine_REDSHIFT EndpointEngine = "REDSHIFT"
	EndpointEngine_KINESIS EndpointEngine = "KINESIS"
	EndpointEngine_KAFKA EndpointEngine = "KAFKA"
	EndpointEngine_OPENSEARCH EndpointEngine = "OPENSEARCH"
	EndpointEngine_NEPTUNE EndpointEngine = "NEPTUNE"
	EndpointEngine_REDIS EndpointEngine = "REDIS"
)

