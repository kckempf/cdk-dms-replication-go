//go:build no_runtime_type_checking

package cdkdmsreplication

// Building without runtime type checking enabled, so all the below just return nil

func validateDmsMigrationPipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDmsMigrationPipelineParameters(scope constructs.Construct, id *string, props *DmsMigrationPipelineProps) error {
	return nil
}

