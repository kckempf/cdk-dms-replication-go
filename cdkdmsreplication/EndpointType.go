package cdkdmsreplication


// Whether an endpoint is a migration source or target.
type EndpointType string

const (
	EndpointType_SOURCE EndpointType = "SOURCE"
	EndpointType_TARGET EndpointType = "TARGET"
)

