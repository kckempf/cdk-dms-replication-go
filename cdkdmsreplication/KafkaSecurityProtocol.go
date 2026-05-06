package cdkdmsreplication


// Security protocol for Kafka endpoints.
type KafkaSecurityProtocol string

const (
	KafkaSecurityProtocol_PLAINTEXT KafkaSecurityProtocol = "PLAINTEXT"
	KafkaSecurityProtocol_SSL KafkaSecurityProtocol = "SSL"
	KafkaSecurityProtocol_SASL_SSL KafkaSecurityProtocol = "SASL_SSL"
	KafkaSecurityProtocol_SSL_AUTHENTICATION KafkaSecurityProtocol = "SSL_AUTHENTICATION"
)

