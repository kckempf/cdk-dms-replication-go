package cdkdmsreplication


// Settings for PostgreSQL and Aurora PostgreSQL endpoints.
type PostgreSqlSettings struct {
	// SQL to run after connecting.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Whether DMS captures DDL events and creates a replication slot.
	CaptureDdls *bool `field:"optional" json:"captureDdls" yaml:"captureDdls"`
	// Schema in which the operational DDL database artifacts are created.
	DdlArtifactsSchema *string `field:"optional" json:"ddlArtifactsSchema" yaml:"ddlArtifactsSchema"`
	// Sets the client statement timeout for the PostgreSQL instance.
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// Whether DMS fails tasks that attempt to truncate a LOB.
	FailTasksOnLobTruncation *bool `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// Whether DMS enables heartbeat signals.
	HeartbeatEnable *bool `field:"optional" json:"heartbeatEnable" yaml:"heartbeatEnable"`
	// The number of seconds between heartbeat signal calls.
	HeartbeatFrequency *float64 `field:"optional" json:"heartbeatFrequency" yaml:"heartbeatFrequency"`
	// Schema to store the heartbeat table.
	HeartbeatSchema *string `field:"optional" json:"heartbeatSchema" yaml:"heartbeatSchema"`
	// When true, maps boolean as boolean instead of varchar(5).
	MapBooleanAsBoolean *bool `field:"optional" json:"mapBooleanAsBoolean" yaml:"mapBooleanAsBoolean"`
	// Maximum file size (in KB) for CSV files created during full load.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// CDC plugin to use.
	PluginName PostgresCdcPlugin `field:"optional" json:"pluginName" yaml:"pluginName"`
	// ARN of the IAM role that provides access to Secrets Manager.
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Full ARN or name of the Secrets Manager secret.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Name of the logical replication slot created for CDC.
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

