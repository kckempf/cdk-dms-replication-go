package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Top-level props for {@link DmsMigrationPipeline}.
type DmsMigrationPipelineProps struct {
	// The type of migration to perform.
	MigrationType MigrationType `field:"required" json:"migrationType" yaml:"migrationType"`
	// VPC in which to place the replication instance.
	//
	// The instance is placed in private subnets.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Allocated storage for the replication instance in GB.
	// Default: 100.
	//
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// CDC start position (LSN or equivalent).
	//
	// Only used when migrationType includes CDC.
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// CDC start time (ISO-8601 string or Unix epoch seconds).
	//
	// Only used when migrationType includes CDC.
	CdcStartTime *string `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// CDC stop position.
	//
	// Only used when migrationType includes CDC.
	CdcStopPosition *string `field:"optional" json:"cdcStopPosition" yaml:"cdcStopPosition"`
	// Whether to create the two account-level DMS service roles (`dms-vpc-role` and `dms-cloudwatch-logs-role`) required by DMS.
	//
	// Set this to `false` if the roles already exist in the AWS account — for
	// example, because another CDK stack (or a manual deployment) already
	// created them. Attempting to create roles with the same name twice in the
	// same account causes a CloudFormation `EntityAlreadyExists` error.
	//
	// When `false`, the construct expects the roles to already be present and
	// skips creating them. The `dmsVpcRole` and `dmsCloudWatchRole` properties
	// will be `undefined`.
	// Default: true.
	//
	CreateDmsServiceRoles *bool `field:"optional" json:"createDmsServiceRoles" yaml:"createDmsServiceRoles"`
	// Whether to create a CloudWatch log group for the task.
	// Default: true.
	//
	EnableCloudWatchLogs *bool `field:"optional" json:"enableCloudWatchLogs" yaml:"enableCloudWatchLogs"`
	// KMS key for encrypting the replication instance storage at rest.
	//
	// A new key is created if not provided.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Engine version for the replication instance.
	// Default: latest version available in the region (chosen by DMS).
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// An existing {@link IDmsEndpoint} to use as the source.
	//
	// Provide this OR `sourceEndpoint` — not both.
	ExistingSourceEndpoint IDmsEndpoint `field:"optional" json:"existingSourceEndpoint" yaml:"existingSourceEndpoint"`
	// An existing {@link IDmsEndpoint} to use as the target.
	//
	// Provide this OR `targetEndpoint` — not both.
	ExistingTargetEndpoint IDmsEndpoint `field:"optional" json:"existingTargetEndpoint" yaml:"existingTargetEndpoint"`
	// Retention period for the CloudWatch log group.
	// Default: logs.RetentionDays.ONE_MONTH
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Whether the replication instance is Multi-AZ.
	// Default: false.
	//
	MultiAz *bool `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Removal policy applied to all resources in the pipeline.
	// Default: cdk.RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Replication instance class.
	// Default: ReplicationInstanceClass.R6I_LARGE
	//
	ReplicationInstanceClass ReplicationInstanceClass `field:"optional" json:"replicationInstanceClass" yaml:"replicationInstanceClass"`
	// Source endpoint configuration.
	//
	// Provide this OR `existingSourceEndpoint` — not both.
	SourceEndpoint *SourceEndpointOptions `field:"optional" json:"sourceEndpoint" yaml:"sourceEndpoint"`
	// Table mappings JSON string.
	//
	// Use {@link TableMappings} to build this.
	// Defaults to "include all tables in all schemas" if not provided.
	TableMappings *string `field:"optional" json:"tableMappings" yaml:"tableMappings"`
	// Target endpoint configuration.
	//
	// Provide this OR `existingTargetEndpoint` — not both.
	TargetEndpoint *TargetEndpointOptions `field:"optional" json:"targetEndpoint" yaml:"targetEndpoint"`
	// Task settings JSON string.
	//
	// Use {@link TaskSettings} to build this.
	// Sensible defaults are applied if not provided.
	TaskSettings *string `field:"optional" json:"taskSettings" yaml:"taskSettings"`
	// Subnet selection for the replication instance.
	// Default: private subnets with egress.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

