package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/internal"
)

// A DMS replication task that ties a replication instance to a source and target endpoint and defines what data to migrate and how.
type DmsReplicationTask interface {
	constructs.Construct
	// The underlying CloudFormation replication task resource.
	CfnReplicationTask() awsdms.CfnReplicationTask
	// The tree node.
	Node() constructs.Node
	// ARN of the replication task.
	ReplicationTaskArn() *string
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

// The jsii proxy struct for DmsReplicationTask
type jsiiProxy_DmsReplicationTask struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DmsReplicationTask) CfnReplicationTask() awsdms.CfnReplicationTask {
	var returns awsdms.CfnReplicationTask
	_jsii_.Get(
		j,
		"cfnReplicationTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationTaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskArn",
		&returns,
	)
	return returns
}


func NewDmsReplicationTask(scope constructs.Construct, id *string, props *DmsReplicationTaskProps) DmsReplicationTask {
	_init_.Initialize()

	if err := validateNewDmsReplicationTaskParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsReplicationTask{}

	_jsii_.Create(
		"cdk-dms-replication.DmsReplicationTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDmsReplicationTask_Override(d DmsReplicationTask, scope constructs.Construct, id *string, props *DmsReplicationTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dms-replication.DmsReplicationTask",
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
func DmsReplicationTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsReplicationTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-dms-replication.DmsReplicationTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationTask) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

