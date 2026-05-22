package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/internal"
)

// An L3 CDK pattern construct that provisions a complete DMS migration pipeline:.
//
// - **Replication instance** — placed in private subnets with KMS encryption
//   and a dedicated security group.
// - **Source endpoint** — supports every engine DMS supports.
// - **Target endpoint** — supports every engine DMS supports.
// - **Replication task** — wires the instance, endpoints, table mappings and
//   task settings together.
// - **IAM role** — grants DMS permission to write to CloudWatch Logs.
// - **CloudWatch log group** — (optional) retains task logs.
//
// Example:
//   // MySQL → Aurora PostgreSQL, full-load-and-CDC
//   const pipeline = new DmsMigrationPipeline(this, 'MigrationPipeline', {
//     vpc,
//     migrationType: MigrationType.FULL_LOAD_AND_CDC,
//     sourceEndpoint: {
//       engine: EndpointEngine.MYSQL,
//       serverName: 'mysql.example.com',
//       port: 3306,
//       username: 'dms_user',
//       password: cdk.SecretValue.secretsManager('mysql-dms-secret'),
//       databaseName: 'mydb',
//     },
//     targetEndpoint: {
//       engine: EndpointEngine.AURORA_POSTGRESQL,
//       serverName: cluster.clusterEndpoint.hostname,
//       port: 5432,
//       username: 'dms_user',
//       password: cdk.SecretValue.secretsManager('aurora-dms-secret'),
//       databaseName: 'mydb',
//     },
//     tableMappings: new TableMappings().includeSchema('public').toJson(),
//   });
//
type DmsMigrationPipeline interface {
	constructs.Construct
	// Construct wrapping the custom resources that created the `dms-cloudwatch-logs-role`.
	//
	// `undefined` when `createDmsServiceRoles` is `false`.
	DmsCloudWatchRole() constructs.Construct
	// Construct wrapping the custom resources that created the `dms-vpc-role`.
	//
	// `undefined` when `createDmsServiceRoles` is `false`.
	DmsVpcRole() constructs.Construct
	// CloudWatch log group for the task (if enableCloudWatchLogs is true).
	LogGroup() awslogs.LogGroup
	// The tree node.
	Node() constructs.Node
	// The replication instance provisioned by this pipeline.
	ReplicationInstance() DmsReplicationInstance
	// The replication task that drives the migration.
	ReplicationTask() DmsReplicationTask
	// The source endpoint used by this pipeline.
	Source() IDmsEndpoint
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

// The jsii proxy struct for DmsMigrationPipeline
type jsiiProxy_DmsMigrationPipeline struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DmsMigrationPipeline) DmsCloudWatchRole() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"dmsCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationPipeline) DmsVpcRole() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"dmsVpcRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationPipeline) LogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationPipeline) ReplicationInstance() DmsReplicationInstance {
	var returns DmsReplicationInstance
	_jsii_.Get(
		j,
		"replicationInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationPipeline) ReplicationTask() DmsReplicationTask {
	var returns DmsReplicationTask
	_jsii_.Get(
		j,
		"replicationTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationPipeline) Source() IDmsEndpoint {
	var returns IDmsEndpoint
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationPipeline) Target() IDmsEndpoint {
	var returns IDmsEndpoint
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


func NewDmsMigrationPipeline(scope constructs.Construct, id *string, props *DmsMigrationPipelineProps) DmsMigrationPipeline {
	_init_.Initialize()

	if err := validateNewDmsMigrationPipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsMigrationPipeline{}

	_jsii_.Create(
		"cdk-dms-replication.DmsMigrationPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDmsMigrationPipeline_Override(d DmsMigrationPipeline, scope constructs.Construct, id *string, props *DmsMigrationPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dms-replication.DmsMigrationPipeline",
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
func DmsMigrationPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsMigrationPipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-dms-replication.DmsMigrationPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationPipeline) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

