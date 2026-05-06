//go:build no_runtime_type_checking

package cdkdmsreplication

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskSettings) validateWithBatchApplyParameters(enabled *bool) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithCommitRateParameters(rows *float64) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithDataErrorPolicyParameters(policy ErrorAction) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithFullLoadSubTasksParameters(count *float64) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithLobModeParameters(mode LobMode) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithLoggingParameters(component *string, level LoggingLevel) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithLoggingEnabledParameters(enabled *bool) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithRecoveryParameters(count *float64) error {
	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithTargetTablePrepModeParameters(mode *string) error {
	return nil
}

