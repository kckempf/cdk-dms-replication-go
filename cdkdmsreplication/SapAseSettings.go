package cdkdmsreplication


// Settings for SAP Adaptive Server Enterprise (Sybase) endpoints.
type SapAseSettings struct {
	// ARN of IAM role for Secrets Manager.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
}

