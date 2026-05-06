package cdkdmsreplication


// MongoDB authentication mechanism.
type MongoAuthMechanism string

const (
	MongoAuthMechanism_DEFAULT MongoAuthMechanism = "DEFAULT"
	MongoAuthMechanism_MONGODB_CR MongoAuthMechanism = "MONGODB_CR"
	MongoAuthMechanism_SCRAM_SHA_1 MongoAuthMechanism = "SCRAM_SHA_1"
)

