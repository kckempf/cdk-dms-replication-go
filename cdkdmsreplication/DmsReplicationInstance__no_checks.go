//go:build no_runtime_type_checking

package cdkdmsreplication

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DmsReplicationInstance) validateAllowInboundFromParameters(peer awsec2.IPeer, port awsec2.Port) error {
	return nil
}

func validateDmsReplicationInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDmsReplicationInstanceParameters(scope constructs.Construct, id *string, props *DmsReplicationInstanceProps) error {
	return nil
}

