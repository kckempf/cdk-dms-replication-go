package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for the target endpoint of a {@link DmsMigrationPipeline}.
//
// Extends {@link DmsEndpointProps} but omits `endpointType` (always TARGET).
type TargetEndpointOptions struct {
	Engine EndpointEngine `field:"required" json:"engine" yaml:"engine"`
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	Db2Settings *Db2Settings `field:"optional" json:"db2Settings" yaml:"db2Settings"`
	DynamoDbSettings *DynamoDbSettings `field:"optional" json:"dynamoDbSettings" yaml:"dynamoDbSettings"`
	EndpointIdentifier *string `field:"optional" json:"endpointIdentifier" yaml:"endpointIdentifier"`
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	KafkaSettings *KafkaSettings `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	KinesisSettings *KinesisSettings `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	MongoDbSettings *MongoDbSettings `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	MySqlSettings *MySqlSettings `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	NeptuneSettings *NeptuneSettings `field:"optional" json:"neptuneSettings" yaml:"neptuneSettings"`
	OpenSearchSettings *OpenSearchSettings `field:"optional" json:"openSearchSettings" yaml:"openSearchSettings"`
	OracleSettings *OracleSettings `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// Database password.
	//
	// The resolved value is stored as **plaintext** in the
	// CloudFormation template. Prefer `secretsManagerSecretId` in the
	// engine-specific settings for production workloads.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	PostgreSqlSettings *PostgreSqlSettings `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	RedisSettings *RedisSettings `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	RedshiftSettings *RedshiftSettings `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	S3Settings *S3Settings `field:"optional" json:"s3Settings" yaml:"s3Settings"`
	SapAseSettings *SapAseSettings `field:"optional" json:"sapAseSettings" yaml:"sapAseSettings"`
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	SqlServerSettings *SqlServerSettings `field:"optional" json:"sqlServerSettings" yaml:"sqlServerSettings"`
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	Username *string `field:"optional" json:"username" yaml:"username"`
}

