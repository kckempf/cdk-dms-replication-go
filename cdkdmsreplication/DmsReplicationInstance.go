package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/internal"
)

// An L2-style construct that provisions a DMS replication instance with: - Private subnet placement via a dedicated subnet group - KMS encryption at rest (key created if not provided) - A dedicated security group (created if not provided).
type DmsReplicationInstance interface {
	constructs.Construct
	// The underlying CloudFormation replication instance resource.
	CfnReplicationInstance() awsdms.CfnReplicationInstance
	// The KMS key used for storage encryption.
	EncryptionKey() awskms.IKey
	// The tree node.
	Node() constructs.Node
	// The replication instance ARN.
	ReplicationInstanceArn() *string
	// The security group attached to this instance.
	SecurityGroup() awsec2.ISecurityGroup
	// The replication subnet group created for this instance.
	SubnetGroup() awsdms.CfnReplicationSubnetGroup
	// Allow inbound access to the replication instance on a given port from a peer.
	//
	// Useful when the source or target database security group needs to trust the instance.
	AllowInboundFrom(peer awsec2.IPeer, port awsec2.Port, description *string)
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

// The jsii proxy struct for DmsReplicationInstance
type jsiiProxy_DmsReplicationInstance struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DmsReplicationInstance) CfnReplicationInstance() awsdms.CfnReplicationInstance {
	var returns awsdms.CfnReplicationInstance
	_jsii_.Get(
		j,
		"cfnReplicationInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) SubnetGroup() awsdms.CfnReplicationSubnetGroup {
	var returns awsdms.CfnReplicationSubnetGroup
	_jsii_.Get(
		j,
		"subnetGroup",
		&returns,
	)
	return returns
}


func NewDmsReplicationInstance(scope constructs.Construct, id *string, props *DmsReplicationInstanceProps) DmsReplicationInstance {
	_init_.Initialize()

	if err := validateNewDmsReplicationInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsReplicationInstance{}

	_jsii_.Create(
		"cdk-dms-replication.DmsReplicationInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDmsReplicationInstance_Override(d DmsReplicationInstance, scope constructs.Construct, id *string, props *DmsReplicationInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dms-replication.DmsReplicationInstance",
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
func DmsReplicationInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsReplicationInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-dms-replication.DmsReplicationInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationInstance) AllowInboundFrom(peer awsec2.IPeer, port awsec2.Port, description *string) {
	if err := d.validateAllowInboundFromParameters(peer, port); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"allowInboundFrom",
		[]interface{}{peer, port, description},
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

