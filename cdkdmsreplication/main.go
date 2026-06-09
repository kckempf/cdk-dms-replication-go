// L3 CDK Constructs for Amazon Database Migration Service (DMS) — pattern construct bundling replication instance, endpoints, and tasks with secure defaults.
package cdkdmsreplication

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk-dms-replication.AddColumnDefinition",
		reflect.TypeOf((*AddColumnDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.CdcSettings",
		reflect.TypeOf((*CdcSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.ColumnDataType",
		reflect.TypeOf((*ColumnDataType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": ColumnDataType_STRING,
			"INT4": ColumnDataType_INT4,
			"INT8": ColumnDataType_INT8,
			"FLOAT4": ColumnDataType_FLOAT4,
			"FLOAT8": ColumnDataType_FLOAT8,
			"NUMERIC": ColumnDataType_NUMERIC,
			"DATETIME": ColumnDataType_DATETIME,
			"BYTES": ColumnDataType_BYTES,
			"BLOB": ColumnDataType_BLOB,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.DatePartitionDelimiter",
		reflect.TypeOf((*DatePartitionDelimiter)(nil)).Elem(),
		map[string]interface{}{
			"SLASH": DatePartitionDelimiter_SLASH,
			"UNDERSCORE": DatePartitionDelimiter_UNDERSCORE,
			"DASH": DatePartitionDelimiter_DASH,
			"NONE": DatePartitionDelimiter_NONE,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.DatePartitionSequence",
		reflect.TypeOf((*DatePartitionSequence)(nil)).Elem(),
		map[string]interface{}{
			"YYYYMMDD": DatePartitionSequence_YYYYMMDD,
			"YYYYMMDDHH": DatePartitionSequence_YYYYMMDDHH,
			"YYYYMM": DatePartitionSequence_YYYYMM,
			"MMYYYYDD": DatePartitionSequence_MMYYYYDD,
			"DDMMYYYY": DatePartitionSequence_DDMMYYYY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.Db2Settings",
		reflect.TypeOf((*Db2Settings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-dms-replication.DmsEndpoint",
		reflect.TypeOf((*DmsEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnEndpoint", GoGetter: "CfnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DmsEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDmsEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DmsEndpointProps",
		reflect.TypeOf((*DmsEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-dms-replication.DmsMigrationPipeline",
		reflect.TypeOf((*DmsMigrationPipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dmsCloudWatchRole", GoGetter: "DmsCloudWatchRole"},
			_jsii_.MemberProperty{JsiiProperty: "dmsVpcRole", GoGetter: "DmsVpcRole"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInstance", GoGetter: "ReplicationInstance"},
			_jsii_.MemberProperty{JsiiProperty: "replicationTask", GoGetter: "ReplicationTask"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DmsMigrationPipeline{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DmsMigrationPipelineProps",
		reflect.TypeOf((*DmsMigrationPipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-dms-replication.DmsReplicationInstance",
		reflect.TypeOf((*DmsReplicationInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allowInboundFrom", GoMethod: "AllowInboundFrom"},
			_jsii_.MemberProperty{JsiiProperty: "cfnReplicationInstance", GoGetter: "CfnReplicationInstance"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInstanceArn", GoGetter: "ReplicationInstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroup", GoGetter: "SubnetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DmsReplicationInstance{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DmsReplicationInstanceProps",
		reflect.TypeOf((*DmsReplicationInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-dms-replication.DmsReplicationTask",
		reflect.TypeOf((*DmsReplicationTask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnReplicationTask", GoGetter: "CfnReplicationTask"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "replicationTaskArn", GoGetter: "ReplicationTaskArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DmsReplicationTask{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DmsReplicationTaskProps",
		reflect.TypeOf((*DmsReplicationTaskProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-dms-replication.DmsServerlessPipeline",
		reflect.TypeOf((*DmsServerlessPipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnReplicationConfig", GoGetter: "CfnReplicationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "dmsCloudWatchRole", GoGetter: "DmsCloudWatchRole"},
			_jsii_.MemberProperty{JsiiProperty: "dmsVpcRole", GoGetter: "DmsVpcRole"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "replicationConfigArn", GoGetter: "ReplicationConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroup", GoGetter: "SubnetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DmsServerlessPipeline{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DmsServerlessPipelineProps",
		reflect.TypeOf((*DmsServerlessPipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DynamoDbAttributeMapping",
		reflect.TypeOf((*DynamoDbAttributeMapping)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.DynamoDbAttributeSubType",
		reflect.TypeOf((*DynamoDbAttributeSubType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": DynamoDbAttributeSubType_STRING,
			"NUMBER": DynamoDbAttributeSubType_NUMBER,
			"BINARY": DynamoDbAttributeSubType_BINARY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DynamoDbKeyMapping",
		reflect.TypeOf((*DynamoDbKeyMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DynamoDbObjectMappingOptions",
		reflect.TypeOf((*DynamoDbObjectMappingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.DynamoDbSettings",
		reflect.TypeOf((*DynamoDbSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.EncodingType",
		reflect.TypeOf((*EncodingType)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": EncodingType_PLAIN,
			"PLAIN_DICTIONARY": EncodingType_PLAIN_DICTIONARY,
			"RLE_DICTIONARY": EncodingType_RLE_DICTIONARY,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.EncryptionMode",
		reflect.TypeOf((*EncryptionMode)(nil)).Elem(),
		map[string]interface{}{
			"SSE_S3": EncryptionMode_SSE_S3,
			"SSE_KMS": EncryptionMode_SSE_KMS,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.EndpointEngine",
		reflect.TypeOf((*EndpointEngine)(nil)).Elem(),
		map[string]interface{}{
			"MYSQL": EndpointEngine_MYSQL,
			"AURORA_MYSQL": EndpointEngine_AURORA_MYSQL,
			"POSTGRES": EndpointEngine_POSTGRES,
			"AURORA_POSTGRESQL": EndpointEngine_AURORA_POSTGRESQL,
			"ORACLE": EndpointEngine_ORACLE,
			"SQLSERVER": EndpointEngine_SQLSERVER,
			"MARIADB": EndpointEngine_MARIADB,
			"SAP_ASE": EndpointEngine_SAP_ASE,
			"IBM_DB2": EndpointEngine_IBM_DB2,
			"IBM_DB2_ZOS": EndpointEngine_IBM_DB2_ZOS,
			"MONGODB": EndpointEngine_MONGODB,
			"DOCDB": EndpointEngine_DOCDB,
			"S3": EndpointEngine_S3,
			"DYNAMODB": EndpointEngine_DYNAMODB,
			"REDSHIFT": EndpointEngine_REDSHIFT,
			"KINESIS": EndpointEngine_KINESIS,
			"KAFKA": EndpointEngine_KAFKA,
			"OPENSEARCH": EndpointEngine_OPENSEARCH,
			"NEPTUNE": EndpointEngine_NEPTUNE,
			"REDIS": EndpointEngine_REDIS,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.EndpointType",
		reflect.TypeOf((*EndpointType)(nil)).Elem(),
		map[string]interface{}{
			"SOURCE": EndpointType_SOURCE,
			"TARGET": EndpointType_TARGET,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.ErrorAction",
		reflect.TypeOf((*ErrorAction)(nil)).Elem(),
		map[string]interface{}{
			"IGNORE": ErrorAction_IGNORE,
			"IGNORE_RECORD": ErrorAction_IGNORE_RECORD,
			"STOP_TASK": ErrorAction_STOP_TASK,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.ErrorHandlingSettings",
		reflect.TypeOf((*ErrorHandlingSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.FullLoadSettings",
		reflect.TypeOf((*FullLoadSettings)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-dms-replication.IDmsEndpoint",
		reflect.TypeOf((*IDmsEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
		},
		func() interface{} {
			return &jsiiProxy_IDmsEndpoint{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.KafkaSecurityProtocol",
		reflect.TypeOf((*KafkaSecurityProtocol)(nil)).Elem(),
		map[string]interface{}{
			"PLAINTEXT": KafkaSecurityProtocol_PLAINTEXT,
			"SSL": KafkaSecurityProtocol_SSL,
			"SASL_SSL": KafkaSecurityProtocol_SASL_SSL,
			"SSL_AUTHENTICATION": KafkaSecurityProtocol_SSL_AUTHENTICATION,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.KafkaSettings",
		reflect.TypeOf((*KafkaSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.KinesisSettings",
		reflect.TypeOf((*KinesisSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.LobMode",
		reflect.TypeOf((*LobMode)(nil)).Elem(),
		map[string]interface{}{
			"NONE": LobMode_NONE,
			"FULL_LOB": LobMode_FULL_LOB,
			"LIMITED_LOB": LobMode_LIMITED_LOB,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.LogComponentSettings",
		reflect.TypeOf((*LogComponentSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.LoggingLevel",
		reflect.TypeOf((*LoggingLevel)(nil)).Elem(),
		map[string]interface{}{
			"LOGGER_SEVERITY_DEFAULT": LoggingLevel_LOGGER_SEVERITY_DEFAULT,
			"LOGGER_SEVERITY_DEBUG": LoggingLevel_LOGGER_SEVERITY_DEBUG,
			"LOGGER_SEVERITY_DETAILED_DEBUG": LoggingLevel_LOGGER_SEVERITY_DETAILED_DEBUG,
			"LOGGER_SEVERITY_ERROR": LoggingLevel_LOGGER_SEVERITY_ERROR,
			"LOGGER_SEVERITY_WARNING": LoggingLevel_LOGGER_SEVERITY_WARNING,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.MessageFormat",
		reflect.TypeOf((*MessageFormat)(nil)).Elem(),
		map[string]interface{}{
			"JSON": MessageFormat_JSON,
			"JSON_UNFORMATTED": MessageFormat_JSON_UNFORMATTED,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.MigrationType",
		reflect.TypeOf((*MigrationType)(nil)).Elem(),
		map[string]interface{}{
			"FULL_LOAD": MigrationType_FULL_LOAD,
			"CDC": MigrationType_CDC,
			"FULL_LOAD_AND_CDC": MigrationType_FULL_LOAD_AND_CDC,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.MongoAuthMechanism",
		reflect.TypeOf((*MongoAuthMechanism)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": MongoAuthMechanism_DEFAULT,
			"MONGODB_CR": MongoAuthMechanism_MONGODB_CR,
			"SCRAM_SHA_1": MongoAuthMechanism_SCRAM_SHA_1,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.MongoAuthType",
		reflect.TypeOf((*MongoAuthType)(nil)).Elem(),
		map[string]interface{}{
			"NO": MongoAuthType_NO,
			"PASSWORD": MongoAuthType_PASSWORD,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.MongoDbSettings",
		reflect.TypeOf((*MongoDbSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.MongoNestingLevel",
		reflect.TypeOf((*MongoNestingLevel)(nil)).Elem(),
		map[string]interface{}{
			"NONE": MongoNestingLevel_NONE,
			"ONE": MongoNestingLevel_ONE,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.MySqlSettings",
		reflect.TypeOf((*MySqlSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.MySqlTargetDbType",
		reflect.TypeOf((*MySqlTargetDbType)(nil)).Elem(),
		map[string]interface{}{
			"SPECIFIC_DATABASE": MySqlTargetDbType_SPECIFIC_DATABASE,
			"MULTIPLE_DATABASES": MySqlTargetDbType_MULTIPLE_DATABASES,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.NeptuneSettings",
		reflect.TypeOf((*NeptuneSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.OpenSearchSettings",
		reflect.TypeOf((*OpenSearchSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.OracleCdcPlugin",
		reflect.TypeOf((*OracleCdcPlugin)(nil)).Elem(),
		map[string]interface{}{
			"LOGMINER": OracleCdcPlugin_LOGMINER,
			"BINARY_READER": OracleCdcPlugin_BINARY_READER,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.OracleNumberDatatypeScale",
		reflect.TypeOf((*OracleNumberDatatypeScale)(nil)).Elem(),
		map[string]interface{}{
			"FLOAT": OracleNumberDatatypeScale_FLOAT,
			"DOUBLE": OracleNumberDatatypeScale_DOUBLE,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.OracleSettings",
		reflect.TypeOf((*OracleSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.ParquetVersion",
		reflect.TypeOf((*ParquetVersion)(nil)).Elem(),
		map[string]interface{}{
			"PARQUET_1_0": ParquetVersion_PARQUET_1_0,
			"PARQUET_2_0": ParquetVersion_PARQUET_2_0,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.PostgreSqlSettings",
		reflect.TypeOf((*PostgreSqlSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.PostgresCdcPlugin",
		reflect.TypeOf((*PostgresCdcPlugin)(nil)).Elem(),
		map[string]interface{}{
			"PG_LOGICAL": PostgresCdcPlugin_PG_LOGICAL,
			"TEST_DECODING": PostgresCdcPlugin_TEST_DECODING,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.RedisSettings",
		reflect.TypeOf((*RedisSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.RedshiftSettings",
		reflect.TypeOf((*RedshiftSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.ReplicationInstanceClass",
		reflect.TypeOf((*ReplicationInstanceClass)(nil)).Elem(),
		map[string]interface{}{
			"T3_SMALL": ReplicationInstanceClass_T3_SMALL,
			"T3_MEDIUM": ReplicationInstanceClass_T3_MEDIUM,
			"T3_LARGE": ReplicationInstanceClass_T3_LARGE,
			"C5_LARGE": ReplicationInstanceClass_C5_LARGE,
			"C5_XLARGE": ReplicationInstanceClass_C5_XLARGE,
			"C5_2XLARGE": ReplicationInstanceClass_C5_2XLARGE,
			"C5_4XLARGE": ReplicationInstanceClass_C5_4XLARGE,
			"C5_9XLARGE": ReplicationInstanceClass_C5_9XLARGE,
			"C5_12XLARGE": ReplicationInstanceClass_C5_12XLARGE,
			"C5_18XLARGE": ReplicationInstanceClass_C5_18XLARGE,
			"C5_24XLARGE": ReplicationInstanceClass_C5_24XLARGE,
			"C6I_LARGE": ReplicationInstanceClass_C6I_LARGE,
			"C6I_XLARGE": ReplicationInstanceClass_C6I_XLARGE,
			"C6I_2XLARGE": ReplicationInstanceClass_C6I_2XLARGE,
			"C6I_4XLARGE": ReplicationInstanceClass_C6I_4XLARGE,
			"C6I_8XLARGE": ReplicationInstanceClass_C6I_8XLARGE,
			"C6I_12XLARGE": ReplicationInstanceClass_C6I_12XLARGE,
			"C6I_16XLARGE": ReplicationInstanceClass_C6I_16XLARGE,
			"C6I_24XLARGE": ReplicationInstanceClass_C6I_24XLARGE,
			"C6I_32XLARGE": ReplicationInstanceClass_C6I_32XLARGE,
			"C7I_LARGE": ReplicationInstanceClass_C7I_LARGE,
			"C7I_XLARGE": ReplicationInstanceClass_C7I_XLARGE,
			"C7I_2XLARGE": ReplicationInstanceClass_C7I_2XLARGE,
			"C7I_4XLARGE": ReplicationInstanceClass_C7I_4XLARGE,
			"C7I_8XLARGE": ReplicationInstanceClass_C7I_8XLARGE,
			"C7I_12XLARGE": ReplicationInstanceClass_C7I_12XLARGE,
			"C7I_16XLARGE": ReplicationInstanceClass_C7I_16XLARGE,
			"C7I_24XLARGE": ReplicationInstanceClass_C7I_24XLARGE,
			"C7I_48XLARGE": ReplicationInstanceClass_C7I_48XLARGE,
			"R5_LARGE": ReplicationInstanceClass_R5_LARGE,
			"R5_XLARGE": ReplicationInstanceClass_R5_XLARGE,
			"R5_2XLARGE": ReplicationInstanceClass_R5_2XLARGE,
			"R5_4XLARGE": ReplicationInstanceClass_R5_4XLARGE,
			"R5_8XLARGE": ReplicationInstanceClass_R5_8XLARGE,
			"R5_12XLARGE": ReplicationInstanceClass_R5_12XLARGE,
			"R5_16XLARGE": ReplicationInstanceClass_R5_16XLARGE,
			"R5_24XLARGE": ReplicationInstanceClass_R5_24XLARGE,
			"R6I_LARGE": ReplicationInstanceClass_R6I_LARGE,
			"R6I_XLARGE": ReplicationInstanceClass_R6I_XLARGE,
			"R6I_2XLARGE": ReplicationInstanceClass_R6I_2XLARGE,
			"R6I_4XLARGE": ReplicationInstanceClass_R6I_4XLARGE,
			"R6I_8XLARGE": ReplicationInstanceClass_R6I_8XLARGE,
			"R6I_12XLARGE": ReplicationInstanceClass_R6I_12XLARGE,
			"R6I_16XLARGE": ReplicationInstanceClass_R6I_16XLARGE,
			"R6I_24XLARGE": ReplicationInstanceClass_R6I_24XLARGE,
			"R6I_32XLARGE": ReplicationInstanceClass_R6I_32XLARGE,
			"R7I_LARGE": ReplicationInstanceClass_R7I_LARGE,
			"R7I_XLARGE": ReplicationInstanceClass_R7I_XLARGE,
			"R7I_2XLARGE": ReplicationInstanceClass_R7I_2XLARGE,
			"R7I_4XLARGE": ReplicationInstanceClass_R7I_4XLARGE,
			"R7I_8XLARGE": ReplicationInstanceClass_R7I_8XLARGE,
			"R7I_12XLARGE": ReplicationInstanceClass_R7I_12XLARGE,
			"R7I_16XLARGE": ReplicationInstanceClass_R7I_16XLARGE,
			"R7I_24XLARGE": ReplicationInstanceClass_R7I_24XLARGE,
			"R7I_48XLARGE": ReplicationInstanceClass_R7I_48XLARGE,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.RuleObjectLocator",
		reflect.TypeOf((*RuleObjectLocator)(nil)).Elem(),
		map[string]interface{}{
			"SCHEMA": RuleObjectLocator_SCHEMA,
			"TABLE": RuleObjectLocator_TABLE,
			"COLUMN": RuleObjectLocator_COLUMN,
			"TABLE_TABLESPACE": RuleObjectLocator_TABLE_TABLESPACE,
			"INDEX_TABLESPACE": RuleObjectLocator_INDEX_TABLESPACE,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.RuleObjectLocatorValue",
		reflect.TypeOf((*RuleObjectLocatorValue)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.S3DataFormat",
		reflect.TypeOf((*S3DataFormat)(nil)).Elem(),
		map[string]interface{}{
			"CSV": S3DataFormat_CSV,
			"PARQUET": S3DataFormat_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.S3Settings",
		reflect.TypeOf((*S3Settings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.SapAseSettings",
		reflect.TypeOf((*SapAseSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.SelectionAction",
		reflect.TypeOf((*SelectionAction)(nil)).Elem(),
		map[string]interface{}{
			"INCLUDE": SelectionAction_INCLUDE,
			"EXCLUDE": SelectionAction_EXCLUDE,
			"EXPLICIT": SelectionAction_EXPLICIT,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.SourceEndpointOptions",
		reflect.TypeOf((*SourceEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.SqlServerSafeguardPolicy",
		reflect.TypeOf((*SqlServerSafeguardPolicy)(nil)).Elem(),
		map[string]interface{}{
			"RELY_ON_SQL_SERVER_REPLICATION_AGENT": SqlServerSafeguardPolicy_RELY_ON_SQL_SERVER_REPLICATION_AGENT,
			"EXCLUSIVE_AUTOMATIC_TRUNCATION": SqlServerSafeguardPolicy_EXCLUSIVE_AUTOMATIC_TRUNCATION,
			"SHARED_AUTOMATIC_TRUNCATION": SqlServerSafeguardPolicy_SHARED_AUTOMATIC_TRUNCATION,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.SqlServerSettings",
		reflect.TypeOf((*SqlServerSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.TableMappingRule",
		reflect.TypeOf((*TableMappingRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-dms-replication.TableMappings",
		reflect.TypeOf((*TableMappings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addColumn", GoMethod: "AddColumn"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefixToSchema", GoMethod: "AddPrefixToSchema"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefixToTable", GoMethod: "AddPrefixToTable"},
			_jsii_.MemberMethod{JsiiMethod: "addSuffixToSchema", GoMethod: "AddSuffixToSchema"},
			_jsii_.MemberMethod{JsiiMethod: "addSuffixToTable", GoMethod: "AddSuffixToTable"},
			_jsii_.MemberMethod{JsiiMethod: "excludeSchema", GoMethod: "ExcludeSchema"},
			_jsii_.MemberMethod{JsiiMethod: "excludeTable", GoMethod: "ExcludeTable"},
			_jsii_.MemberMethod{JsiiMethod: "explicitTable", GoMethod: "ExplicitTable"},
			_jsii_.MemberMethod{JsiiMethod: "includeSchema", GoMethod: "IncludeSchema"},
			_jsii_.MemberMethod{JsiiMethod: "includeTable", GoMethod: "IncludeTable"},
			_jsii_.MemberMethod{JsiiMethod: "mapToDynamoDb", GoMethod: "MapToDynamoDb"},
			_jsii_.MemberMethod{JsiiMethod: "removeColumn", GoMethod: "RemoveColumn"},
			_jsii_.MemberMethod{JsiiMethod: "renameColumn", GoMethod: "RenameColumn"},
			_jsii_.MemberMethod{JsiiMethod: "renameSchema", GoMethod: "RenameSchema"},
			_jsii_.MemberMethod{JsiiMethod: "renameTable", GoMethod: "RenameTable"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toLowerCaseColumn", GoMethod: "ToLowerCaseColumn"},
			_jsii_.MemberMethod{JsiiMethod: "toLowerCaseSchema", GoMethod: "ToLowerCaseSchema"},
			_jsii_.MemberMethod{JsiiMethod: "toLowerCaseTable", GoMethod: "ToLowerCaseTable"},
			_jsii_.MemberMethod{JsiiMethod: "toUpperCaseColumn", GoMethod: "ToUpperCaseColumn"},
			_jsii_.MemberMethod{JsiiMethod: "toUpperCaseSchema", GoMethod: "ToUpperCaseSchema"},
			_jsii_.MemberMethod{JsiiMethod: "toUpperCaseTable", GoMethod: "ToUpperCaseTable"},
		},
		func() interface{} {
			return &jsiiProxy_TableMappings{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dms-replication.TargetEndpointOptions",
		reflect.TypeOf((*TargetEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-dms-replication.TaskSettings",
		reflect.TypeOf((*TaskSettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "withBatchApply", GoMethod: "WithBatchApply"},
			_jsii_.MemberMethod{JsiiMethod: "withCommitRate", GoMethod: "WithCommitRate"},
			_jsii_.MemberMethod{JsiiMethod: "withDataErrorPolicy", GoMethod: "WithDataErrorPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "withFullLoadSubTasks", GoMethod: "WithFullLoadSubTasks"},
			_jsii_.MemberMethod{JsiiMethod: "withLobMode", GoMethod: "WithLobMode"},
			_jsii_.MemberMethod{JsiiMethod: "withLogging", GoMethod: "WithLogging"},
			_jsii_.MemberMethod{JsiiMethod: "withLoggingEnabled", GoMethod: "WithLoggingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "withRecovery", GoMethod: "WithRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "withTargetTablePrepMode", GoMethod: "WithTargetTablePrepMode"},
		},
		func() interface{} {
			return &jsiiProxy_TaskSettings{}
		},
	)
	_jsii_.RegisterEnum(
		"cdk-dms-replication.TransformationAction",
		reflect.TypeOf((*TransformationAction)(nil)).Elem(),
		map[string]interface{}{
			"CONVERT_LOWERCASE": TransformationAction_CONVERT_LOWERCASE,
			"CONVERT_UPPERCASE": TransformationAction_CONVERT_UPPERCASE,
			"ADD_PREFIX": TransformationAction_ADD_PREFIX,
			"REMOVE_PREFIX": TransformationAction_REMOVE_PREFIX,
			"ADD_SUFFIX": TransformationAction_ADD_SUFFIX,
			"REMOVE_SUFFIX": TransformationAction_REMOVE_SUFFIX,
			"RENAME": TransformationAction_RENAME,
			"REMOVE_COLUMN": TransformationAction_REMOVE_COLUMN,
			"ADD_COLUMN": TransformationAction_ADD_COLUMN,
			"INCLUDE_COLUMN": TransformationAction_INCLUDE_COLUMN,
		},
	)
}
