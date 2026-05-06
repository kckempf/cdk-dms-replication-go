package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Props for {@link DmsReplicationInstance}.
type DmsReplicationInstanceProps struct {
	// VPC in which to place the replication instance.
	//
	// The construct creates a dedicated subnet group from the private subnets.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Allocated storage in GB.
	// Default: 100.
	//
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Whether to allow minor version upgrades to be applied automatically during maintenance.
	// Default: true.
	//
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Replication engine version.
	// Default: latest version available in the region (chosen by DMS).
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// KMS key used to encrypt the replication instance storage at rest.
	//
	// A new key is created if not provided.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Whether the replication instance is Multi-AZ.
	// Default: false.
	//
	MultiAz *bool `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Preferred maintenance window, e.g. "sun:04:00-sun:04:30".
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Whether the replication instance is publicly accessible.
	//
	// Setting this to true is strongly discouraged for production workloads.
	// Default: false.
	//
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Removal policy applied to the replication instance.
	// Default: cdk.RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Instance class for the replication instance.
	// Default: ReplicationInstanceClass.R6I_LARGE
	//
	ReplicationInstanceClass ReplicationInstanceClass `field:"optional" json:"replicationInstanceClass" yaml:"replicationInstanceClass"`
	// Logical name of the replication instance.
	//
	// Used in the replication instance identifier and resource naming.
	// Default: id of the construct.
	//
	ReplicationInstanceIdentifier *string `field:"optional" json:"replicationInstanceIdentifier" yaml:"replicationInstanceIdentifier"`
	// Security groups to attach to the replication instance.
	//
	// A new security group is created if none are provided.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specific subnets to use for the replication subnet group.
	//
	// Defaults to all private subnets in the VPC.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

