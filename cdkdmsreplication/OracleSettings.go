package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Settings for Oracle endpoints.
type OracleSettings struct {
	// Set to true to access the supplemental log directly.
	AccessAlternateDirectly *bool `field:"optional" json:"accessAlternateDirectly" yaml:"accessAlternateDirectly"`
	// Additional archived redo log destination ID.
	AdditionalArchivedLogDestId *float64 `field:"optional" json:"additionalArchivedLogDestId" yaml:"additionalArchivedLogDestId"`
	// Whether DMS adds supplemental logging.
	AddSupplementalLogging *bool `field:"optional" json:"addSupplementalLogging" yaml:"addSupplementalLogging"`
	// Allows DMS to access SELECT on nested tables.
	AllowSelectNestedTables *bool `field:"optional" json:"allowSelectNestedTables" yaml:"allowSelectNestedTables"`
	// Destination ID of the archived redo log.
	ArchivedLogDestId *float64 `field:"optional" json:"archivedLogDestId" yaml:"archivedLogDestId"`
	// Reads changes only from archived redo logs (no online redo).
	ArchivedLogsOnly *bool `field:"optional" json:"archivedLogsOnly" yaml:"archivedLogsOnly"`
	// ASM password.
	AsmPassword awscdk.SecretValue `field:"optional" json:"asmPassword" yaml:"asmPassword"`
	// ASM server address.
	AsmServer *string `field:"optional" json:"asmServer" yaml:"asmServer"`
	// ASM user name.
	AsmUser *string `field:"optional" json:"asmUser" yaml:"asmUser"`
	// Semantics for char length: BYTE or CHAR.
	CharLengthSemantics *string `field:"optional" json:"charLengthSemantics" yaml:"charLengthSemantics"`
	// Whether DMS uses direct path full load.
	DirectPathNoLog *bool `field:"optional" json:"directPathNoLog" yaml:"directPathNoLog"`
	// Whether to load in parallel using direct path.
	DirectPathParallelLoad *bool `field:"optional" json:"directPathParallelLoad" yaml:"directPathParallelLoad"`
	// Specifies whether Oracle homogeneous tablespace migration is enabled.
	EnableHomogenousTablespace *bool `field:"optional" json:"enableHomogenousTablespace" yaml:"enableHomogenousTablespace"`
	// Extra archived log destination IDs (up to 10).
	ExtraArchivedLogDestIds *[]*float64 `field:"optional" json:"extraArchivedLogDestIds" yaml:"extraArchivedLogDestIds"`
	// Whether tasks fail if a LOB column is truncated.
	FailTasksOnLobTruncation *bool `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// Precision to use when converting Oracle NUMBER to Amazon Redshift NUMERIC.
	NumberDatatypeScale *float64 `field:"optional" json:"numberDatatypeScale" yaml:"numberDatatypeScale"`
	// Path prefix used for the location of the online redo log.
	OraclePathPrefix *string `field:"optional" json:"oraclePathPrefix" yaml:"oraclePathPrefix"`
	// Number of threads for parallel ASM reading.
	ParallelAsmReadThreads *float64 `field:"optional" json:"parallelAsmReadThreads" yaml:"parallelAsmReadThreads"`
	// Number of read-ahead blocks.
	ReadAheadBlocks *float64 `field:"optional" json:"readAheadBlocks" yaml:"readAheadBlocks"`
	// Read the tablespace name from the online redo log.
	ReadTableSpaceName *bool `field:"optional" json:"readTableSpaceName" yaml:"readTableSpaceName"`
	// Retry interval in seconds when no archivelog is found.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// ARN of IAM role for accessing Secrets Manager (main secret).
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// ARN of IAM role for accessing the ASM Secrets Manager secret.
	SecretsManagerOracleAsmAccessRoleArn *string `field:"optional" json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Full ARN of the ASM secret in Secrets Manager.
	SecretsManagerOracleAsmSecretId *string `field:"optional" json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// Full ARN of the main Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Name of the transparent data encryption (TDE) wallet.
	SecurityDbEncryptionName *string `field:"optional" json:"securityDbEncryptionName" yaml:"securityDbEncryptionName"`
	// Use an alternate folder for online redo logs.
	UseAlternateFolderForOnline *bool `field:"optional" json:"useAlternateFolderForOnline" yaml:"useAlternateFolderForOnline"`
	// Use B-file for large object replication.
	UseBFile *bool `field:"optional" json:"useBFile" yaml:"useBFile"`
	// Use direct path full load.
	UseDirectPathFullLoad *bool `field:"optional" json:"useDirectPathFullLoad" yaml:"useDirectPathFullLoad"`
	// Use LogMiner for CDC.
	UseLogminerReader *bool `field:"optional" json:"useLogminerReader" yaml:"useLogminerReader"`
	// Use a path prefix for the location of the online redo log.
	UsePathPrefix *string `field:"optional" json:"usePathPrefix" yaml:"usePathPrefix"`
}

