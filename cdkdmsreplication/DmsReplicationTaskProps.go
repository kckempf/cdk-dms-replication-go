package cdkdmsreplication

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for {@link DmsReplicationTask}.
type DmsReplicationTaskProps struct {
	// The migration type.
	MigrationType MigrationType `field:"required" json:"migrationType" yaml:"migrationType"`
	// ARN of the replication instance to use.
	ReplicationInstanceArn *string `field:"required" json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
	// Source endpoint.
	SourceEndpoint IDmsEndpoint `field:"required" json:"sourceEndpoint" yaml:"sourceEndpoint"`
	// Table mappings JSON string.
	//
	// Use {@link TableMappings} to build this, or provide a raw JSON string.
	//
	// Example:
	//   tableMappings: new TableMappings().includeSchema('public').toJson()
	//
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// Target endpoint.
	TargetEndpoint IDmsEndpoint `field:"required" json:"targetEndpoint" yaml:"targetEndpoint"`
	// CDC start position (log sequence number or similar).
	//
	// Only valid when migrationType is CDC or FULL_LOAD_AND_CDC.
	// Mutually exclusive with cdcStartTime.
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// CDC start time as an ISO-8601 string or Unix epoch number.
	//
	// Only valid when migrationType is CDC or FULL_LOAD_AND_CDC.
	// Mutually exclusive with cdcStartPosition.
	CdcStartTime *string `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// CDC stop position.
	//
	// The task stops once this position is reached.
	// Only valid when migrationType is CDC or FULL_LOAD_AND_CDC.
	CdcStopPosition *string `field:"optional" json:"cdcStopPosition" yaml:"cdcStopPosition"`
	// Removal policy for this resource.
	// Default: cdk.RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Logical name for the replication task.
	//
	// Auto-generated from the construct ID if not provided.
	ReplicationTaskIdentifier *string `field:"optional" json:"replicationTaskIdentifier" yaml:"replicationTaskIdentifier"`
	// Tags to apply to the task resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Task settings JSON string.
	//
	// Use {@link TaskSettings} to build this, or provide a raw JSON string.
	// A sensible default is applied if not provided.
	TaskSettings *string `field:"optional" json:"taskSettings" yaml:"taskSettings"`
}

