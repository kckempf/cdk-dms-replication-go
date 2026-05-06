//go:build no_runtime_type_checking

package cdkdmsreplication

// Building without runtime type checking enabled, so all the below just return nil

func validateDmsEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDmsEndpointParameters(scope constructs.Construct, id *string, props *DmsEndpointProps) error {
	return nil
}

