package cdkdmsreplication


// Settings shared by MongoDB and DocumentDB endpoints.
type MongoDbSettings struct {
	// Authentication mechanism for the MongoDB endpoint.
	AuthMechanism MongoAuthMechanism `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// Database that MongoDB uses to authenticate.
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// Authentication type.
	AuthType MongoAuthType `field:"optional" json:"authType" yaml:"authType"`
	// Number of documents to preview to determine data structure.
	DocsToInvestigate *float64 `field:"optional" json:"docsToInvestigate" yaml:"docsToInvestigate"`
	// Specifies the document ID, which DMS includes as the primary key.
	ExtractDocId *bool `field:"optional" json:"extractDocId" yaml:"extractDocId"`
	// Nesting level for MongoDB documents.
	NestingLevel MongoNestingLevel `field:"optional" json:"nestingLevel" yaml:"nestingLevel"`
	// ARN of IAM role for Secrets Manager.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

