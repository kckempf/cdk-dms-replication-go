//go:build no_runtime_type_checking

package cdkdmsreplication

// Building without runtime type checking enabled, so all the below just return nil

func validateDmsServerlessPipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDmsServerlessPipelineParameters(scope constructs.Construct, id *string, props *DmsServerlessPipelineProps) error {
	return nil
}

