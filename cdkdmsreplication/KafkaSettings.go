package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Settings for Apache Kafka (and Amazon MSK) target endpoints.
type KafkaSettings struct {
	// Kafka broker(s), e.g. "b-1.msk-cluster.abc123.kafka.us-east-1.amazonaws.com:9092".
	Broker *string `field:"required" json:"broker" yaml:"broker"`
	// Whether to include control details in messages.
	IncludeControlDetails *bool `field:"optional" json:"includeControlDetails" yaml:"includeControlDetails"`
	// Whether to include null and empty columns in messages.
	IncludeNullAndEmpty *bool `field:"optional" json:"includeNullAndEmpty" yaml:"includeNullAndEmpty"`
	// Whether to include the partition value in messages.
	IncludePartitionValue *bool `field:"optional" json:"includePartitionValue" yaml:"includePartitionValue"`
	// Whether to include table ALTER operations.
	IncludeTableAlterOperations *bool `field:"optional" json:"includeTableAlterOperations" yaml:"includeTableAlterOperations"`
	// Whether to include transaction details.
	IncludeTransactionDetails *bool `field:"optional" json:"includeTransactionDetails" yaml:"includeTransactionDetails"`
	// Message format.
	MessageFormat MessageFormat `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Maximum message size in bytes.
	MessageMaxBytes *float64 `field:"optional" json:"messageMaxBytes" yaml:"messageMaxBytes"`
	// Whether to omit the hex prefix from binary values.
	NoHexPrefix *bool `field:"optional" json:"noHexPrefix" yaml:"noHexPrefix"`
	// Whether to include the schema name in the partition key.
	PartitionIncludeSchemaTable *bool `field:"optional" json:"partitionIncludeSchemaTable" yaml:"partitionIncludeSchemaTable"`
	// SASL password.
	SaslPassword awscdk.SecretValue `field:"optional" json:"saslPassword" yaml:"saslPassword"`
	// SASL username.
	SaslUsername *string `field:"optional" json:"saslUsername" yaml:"saslUsername"`
	// Security protocol for the Kafka connection.
	SecurityProtocol KafkaSecurityProtocol `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// ARN of the Secrets Manager secret holding the CA certificate.
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// ARN of the Secrets Manager secret holding the client certificate.
	SslClientCertificateArn *string `field:"optional" json:"sslClientCertificateArn" yaml:"sslClientCertificateArn"`
	// ARN of the Secrets Manager secret holding the client key.
	SslClientKeyArn *string `field:"optional" json:"sslClientKeyArn" yaml:"sslClientKeyArn"`
	// Password for the client key.
	SslClientKeyPassword awscdk.SecretValue `field:"optional" json:"sslClientKeyPassword" yaml:"sslClientKeyPassword"`
	// Topic name to publish to.
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

