package cdkdmsreplication


// Settings for Microsoft SQL Server endpoints.
type SqlServerSettings struct {
	// Maximum number of bytes per BCP packet.
	BcpPacketSize *float64 `field:"optional" json:"bcpPacketSize" yaml:"bcpPacketSize"`
	// Filegroup in SQL Server for control tables.
	ControlTablesFileGroup *string `field:"optional" json:"controlTablesFileGroup" yaml:"controlTablesFileGroup"`
	// Whether to query a single AlwaysOn node.
	QuerySingleAlwaysOnNode *bool `field:"optional" json:"querySingleAlwaysOnNode" yaml:"querySingleAlwaysOnNode"`
	// Whether to use backup files for CDC.
	ReadBackupOnly *bool `field:"optional" json:"readBackupOnly" yaml:"readBackupOnly"`
	// Safeguard policy for SQL Server CDC.
	SafeguardPolicy SqlServerSafeguardPolicy `field:"optional" json:"safeguardPolicy" yaml:"safeguardPolicy"`
	// ARN of IAM role for Secrets Manager.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Controls how DMS accesses the SQL Server transaction log.
	TlogAccessMode *string `field:"optional" json:"tlogAccessMode" yaml:"tlogAccessMode"`
	// Trim spaces from CHAR/VARCHAR columns.
	TrimSpaceInChar *bool `field:"optional" json:"trimSpaceInChar" yaml:"trimSpaceInChar"`
	// Use BCP full load.
	UseBcpFullLoad *bool `field:"optional" json:"useBcpFullLoad" yaml:"useBcpFullLoad"`
	// Use third-party backup device.
	UseThirdPartyBackupDevice *bool `field:"optional" json:"useThirdPartyBackupDevice" yaml:"useThirdPartyBackupDevice"`
}

