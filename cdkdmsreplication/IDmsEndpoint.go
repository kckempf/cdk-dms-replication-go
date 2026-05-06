package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Minimal contract for a DMS endpoint that can be used as a source or target in a {@link DmsReplicationTask } or {@link DmsMigrationPipeline }.
//
// Use this interface when referencing an endpoint created outside of this
// construct (e.g. by ARN). For endpoints created by this library, use the
// concrete {@link DmsEndpoint} class directly — it exposes `cfnEndpoint`
// for L1 escape-hatch access.
type IDmsEndpoint interface {
	// ARN of the DMS endpoint.
	EndpointArn() *string
}

// The jsii proxy for IDmsEndpoint
type jsiiProxy_IDmsEndpoint struct {
	_ byte // padding
}

func (j *jsiiProxy_IDmsEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

