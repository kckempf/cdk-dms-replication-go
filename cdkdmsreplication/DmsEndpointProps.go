package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for {@link DmsEndpoint}.
type DmsEndpointProps struct {
	// Whether this is a source or target endpoint.
	EndpointType EndpointType `field:"required" json:"endpointType" yaml:"endpointType"`
	// Database engine for this endpoint.
	Engine EndpointEngine `field:"required" json:"engine" yaml:"engine"`
	// ARN of the SSL certificate authority certificate.
	//
	// Required when sslMode is 'verify-ca' or 'verify-full'.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Database name on the endpoint.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Settings for IBM Db2 LUW endpoints.
	Db2Settings *Db2Settings `field:"optional" json:"db2Settings" yaml:"db2Settings"`
	// Settings for Amazon DynamoDB target endpoints.
	DynamoDbSettings *DynamoDbSettings `field:"optional" json:"dynamoDbSettings" yaml:"dynamoDbSettings"`
	// Logical identifier of the endpoint (used as the DMS endpoint identifier).
	//
	// Auto-generated from the construct ID if not provided.
	EndpointIdentifier *string `field:"optional" json:"endpointIdentifier" yaml:"endpointIdentifier"`
	// Extra connection attributes as a semicolon-separated string.
	//
	// Refer to the DMS documentation for engine-specific attributes.
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Settings for Apache Kafka (and Amazon MSK) target endpoints.
	KafkaSettings *KafkaSettings `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// Settings for Amazon Kinesis Data Streams target endpoints.
	KinesisSettings *KinesisSettings `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	// Settings for MongoDB or Amazon DocumentDB endpoints.
	MongoDbSettings *MongoDbSettings `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// Settings for MySQL or MariaDB endpoints.
	MySqlSettings *MySqlSettings `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// Settings for Amazon Neptune target endpoints.
	NeptuneSettings *NeptuneSettings `field:"optional" json:"neptuneSettings" yaml:"neptuneSettings"`
	// Settings for Amazon OpenSearch Service target endpoints.
	OpenSearchSettings *OpenSearchSettings `field:"optional" json:"openSearchSettings" yaml:"openSearchSettings"`
	// Settings for Oracle endpoints.
	OracleSettings *OracleSettings `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// Database password.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
	// Database port.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Settings for PostgreSQL or Aurora PostgreSQL endpoints.
	PostgreSqlSettings *PostgreSqlSettings `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// Settings for Amazon ElastiCache for Redis target endpoints.
	RedisSettings *RedisSettings `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	// Settings for Amazon Redshift target endpoints.
	RedshiftSettings *RedshiftSettings `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// Removal policy for the endpoint resource.
	// Default: cdk.RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Settings for Amazon S3 endpoints.
	S3Settings *S3Settings `field:"optional" json:"s3Settings" yaml:"s3Settings"`
	// Settings for SAP ASE (Sybase) endpoints.
	SapAseSettings *SapAseSettings `field:"optional" json:"sapAseSettings" yaml:"sapAseSettings"`
	// Database server hostname or IP address.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Settings for Microsoft SQL Server endpoints.
	SqlServerSettings *SqlServerSettings `field:"optional" json:"sqlServerSettings" yaml:"sqlServerSettings"`
	// SSL mode for the connection.
	// Default: "none".
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Database user name.
	//
	// Used when not using Secrets Manager.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

