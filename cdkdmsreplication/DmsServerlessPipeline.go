package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/internal"
)

// An L3 CDK pattern construct that provisions a DMS Serverless replication pipeline:.
//
// - **Replication config** — backed by `CfnReplicationConfig`; DMS auto-scales
//   capacity between `minCapacityUnits` and `maxCapacityUnits`.
// - **Source endpoint** — supports every engine DMS supports.
// - **Target endpoint** — supports every engine DMS supports.
// - **Subnet group** — private subnet placement.
// - **KMS key** — storage encryption at rest (created if not provided).
// - **Security group** — dedicated group (created if not provided).
// - **IAM roles** — `dms-vpc-role` and `dms-cloudwatch-logs-role`.
// - **CloudWatch log group** — (optional) retains replication logs.
//
// > **CDC start/stop position limitation**: `CfnReplicationConfig` does not
// > expose `cdcStartTime` / `cdcStartPosition` / `cdcStopPosition`. To start
// > from a specific position, call the `StartReplication` API directly after
// > the config is created (e.g. from a Lambda custom resource or CLI).
//
// Example:
//   const pipeline = new DmsServerlessPipeline(this, 'ServerlessPipeline', {
//     vpc,
//     maxCapacityUnits: 8,
//     migrationType: MigrationType.FULL_LOAD_AND_CDC,
//     sourceEndpoint: {
//       engine: EndpointEngine.MYSQL,
//       serverName: 'mysql.example.com',
//       port: 3306,
//       username: 'dms_user',
//       password: cdk.SecretValue.secretsManager('mysql-dms-secret'),
//     },
//     targetEndpoint: {
//       engine: EndpointEngine.AURORA_POSTGRESQL,
//       serverName: cluster.clusterEndpoint.hostname,
//       port: 5432,
//       username: 'dms_user',
//       password: cdk.SecretValue.secretsManager('aurora-dms-secret'),
//     },
//   });
//
type DmsServerlessPipeline interface {
	constructs.Construct
	// The underlying `CfnReplicationConfig` resource.
	CfnReplicationConfig() awsdms.CfnReplicationConfig
	// IAM role that allows DMS to write to CloudWatch Logs.
	//
	// `undefined` when `createDmsServiceRoles` is `false`.
	DmsCloudWatchRole() awsiam.Role
	// IAM role that allows DMS to manage VPC resources (dms-vpc-role).
	//
	// `undefined` when `createDmsServiceRoles` is `false`.
	DmsVpcRole() awsiam.Role
	// The KMS key used for storage encryption.
	EncryptionKey() awskms.IKey
	// CloudWatch log group for the replication config (if enableCloudWatchLogs is true).
	LogGroup() awslogs.LogGroup
	// The tree node.
	Node() constructs.Node
	// The ARN of the replication config.
	ReplicationConfigArn() *string
	// The security group attached to this pipeline.
	SecurityGroup() awsec2.ISecurityGroup
	// The source endpoint used by this pipeline.
	Source() IDmsEndpoint
	// The replication subnet group created for this pipeline.
	SubnetGroup() awsdms.CfnReplicationSubnetGroup
	// The target endpoint used by this pipeline.
	Target() IDmsEndpoint
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DmsServerlessPipeline
type jsiiProxy_DmsServerlessPipeline struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DmsServerlessPipeline) CfnReplicationConfig() awsdms.CfnReplicationConfig {
	var returns awsdms.CfnReplicationConfig
	_jsii_.Get(
		j,
		"cfnReplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) DmsCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"dmsCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) DmsVpcRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"dmsVpcRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) LogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) ReplicationConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) Source() IDmsEndpoint {
	var returns IDmsEndpoint
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) SubnetGroup() awsdms.CfnReplicationSubnetGroup {
	var returns awsdms.CfnReplicationSubnetGroup
	_jsii_.Get(
		j,
		"subnetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsServerlessPipeline) Target() IDmsEndpoint {
	var returns IDmsEndpoint
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


func NewDmsServerlessPipeline(scope constructs.Construct, id *string, props *DmsServerlessPipelineProps) DmsServerlessPipeline {
	_init_.Initialize()

	if err := validateNewDmsServerlessPipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsServerlessPipeline{}

	_jsii_.Create(
		"cdk-dms-replication.DmsServerlessPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDmsServerlessPipeline_Override(d DmsServerlessPipeline, scope constructs.Construct, id *string, props *DmsServerlessPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dms-replication.DmsServerlessPipeline",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DmsServerlessPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsServerlessPipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-dms-replication.DmsServerlessPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsServerlessPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsServerlessPipeline) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

