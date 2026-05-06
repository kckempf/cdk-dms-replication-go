package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Settings for Amazon ElastiCache for Redis target endpoints.
type RedisSettings struct {
	// Redis server name or cluster endpoint.
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// Auth token for auth-token authentication.
	AuthPassword awscdk.SecretValue `field:"optional" json:"authPassword" yaml:"authPassword"`
	// Authentication type: none, auth-token, or auth-role.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// IAM role ARN for auth-role authentication.
	AuthUserName *string `field:"optional" json:"authUserName" yaml:"authUserName"`
	// Redis port.
	//
	// Default: 6379.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// ARN of the SSL CA certificate stored in Secrets Manager.
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// SSL security protocol.
	SslSecurityProtocol *string `field:"optional" json:"sslSecurityProtocol" yaml:"sslSecurityProtocol"`
}

