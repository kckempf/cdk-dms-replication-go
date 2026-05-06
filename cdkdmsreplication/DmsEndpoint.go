package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/internal"
)

// A DMS endpoint construct supporting every engine that DMS supports.
//
// Set `engine` to the desired {@link EndpointEngine} and supply the
// matching `*Settings` property. All other `*Settings` properties are
// ignored at runtime.
//
// Example:
//   // MySQL source
//   new DmsEndpoint(this, 'Source', {
//     endpointType: EndpointType.SOURCE,
//     engine: EndpointEngine.MYSQL,
//     serverName: 'mysql.example.com',
//     port: 3306,
//     username: 'dms_user',
//     password: cdk.SecretValue.ssmSecure('/dms/mysql/password'),
//     databaseName: 'mydb',
//   });
//
type DmsEndpoint interface {
	constructs.Construct
	IDmsEndpoint
	// The underlying CloudFormation endpoint resource.
	CfnEndpoint() awsdms.CfnEndpoint
	// ARN of the DMS endpoint.
	EndpointArn() *string
	// The tree node.
	Node() constructs.Node
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

// The jsii proxy struct for DmsEndpoint
type jsiiProxy_DmsEndpoint struct {
	internal.Type__constructsConstruct
	jsiiProxy_IDmsEndpoint
}

func (j *jsiiProxy_DmsEndpoint) CfnEndpoint() awsdms.CfnEndpoint {
	var returns awsdms.CfnEndpoint
	_jsii_.Get(
		j,
		"cfnEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewDmsEndpoint(scope constructs.Construct, id *string, props *DmsEndpointProps) DmsEndpoint {
	_init_.Initialize()

	if err := validateNewDmsEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpoint{}

	_jsii_.Create(
		"cdk-dms-replication.DmsEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDmsEndpoint_Override(d DmsEndpoint, scope constructs.Construct, id *string, props *DmsEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dms-replication.DmsEndpoint",
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
func DmsEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-dms-replication.DmsEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

