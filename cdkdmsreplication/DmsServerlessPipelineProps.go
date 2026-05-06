package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Props for {@link DmsServerlessPipeline}.
type DmsServerlessPipelineProps struct {
	// Maximum number of DMS Capacity Units (DCUs) the serverless replication may scale up to.
	//
	// Valid values: 1, 2, 4, 8, 16, 32, 64, 128, 192, 256, 384.
	MaxCapacityUnits *float64 `field:"required" json:"maxCapacityUnits" yaml:"maxCapacityUnits"`
	// The type of migration to perform.
	MigrationType MigrationType `field:"required" json:"migrationType" yaml:"migrationType"`
	// VPC in which to place the serverless replication config.
	//
	// The config is placed in private subnets.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to create the two account-level DMS service roles (`dms-vpc-role` and `dms-cloudwatch-logs-role`) required by DMS.
	//
	// Set to `false` if the roles already exist — for example, because a
	// `DmsMigrationPipeline` or a prior manual deployment already created them.
	// Attempting to create roles with the same name twice causes a CloudFormation
	// `EntityAlreadyExists` error.
	// Default: true.
	//
	CreateDmsServiceRoles *bool `field:"optional" json:"createDmsServiceRoles" yaml:"createDmsServiceRoles"`
	// Whether to create a CloudWatch log group for the replication config.
	// Default: true.
	//
	EnableCloudWatchLogs *bool `field:"optional" json:"enableCloudWatchLogs" yaml:"enableCloudWatchLogs"`
	// An existing {@link IDmsEndpoint} to use as the source.
	//
	// Provide this OR `sourceEndpoint` — not both.
	ExistingSourceEndpoint IDmsEndpoint `field:"optional" json:"existingSourceEndpoint" yaml:"existingSourceEndpoint"`
	// An existing {@link IDmsEndpoint} to use as the target.
	//
	// Provide this OR `targetEndpoint` — not both.
	ExistingTargetEndpoint IDmsEndpoint `field:"optional" json:"existingTargetEndpoint" yaml:"existingTargetEndpoint"`
	// KMS key used to encrypt replication storage at rest.
	//
	// A new key is created if not provided.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Retention period for the CloudWatch log group.
	// Default: logs.RetentionDays.ONE_MONTH
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Minimum number of DCUs DMS will provision at start-up.
	// Default: DMS auto-determines based on workload.
	//
	MinCapacityUnits *float64 `field:"optional" json:"minCapacityUnits" yaml:"minCapacityUnits"`
	// Whether the serverless replication is Multi-AZ.
	// Default: false.
	//
	MultiAz *bool `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Preferred maintenance window, e.g. "sun:04:00-sun:04:30".
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Removal policy applied to all resources in the pipeline.
	// Default: cdk.RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Logical identifier for the replication config resource.
	// Default: unique name derived from the construct id.
	//
	ReplicationConfigIdentifier *string `field:"optional" json:"replicationConfigIdentifier" yaml:"replicationConfigIdentifier"`
	// Security groups to attach to the serverless replication config.
	//
	// A new security group is created if none are provided.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
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
	// Replication settings JSON string (equivalent to task settings for classic DMS).
	//
	// Use {@link TaskSettings } to build this.
	// Sensible defaults are applied by DMS if not provided.
	TaskSettings *string `field:"optional" json:"taskSettings" yaml:"taskSettings"`
	// Subnet selection for the replication subnet group.
	// Default: private subnets with egress.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

