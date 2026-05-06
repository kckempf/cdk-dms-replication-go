//go:build !no_runtime_type_checking

package cdkdmsreplication

import (
	"fmt"
)

func (t *jsiiProxy_TaskSettings) validateWithBatchApplyParameters(enabled *bool) error {
	if enabled == nil {
		return fmt.Errorf("parameter enabled is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithCommitRateParameters(rows *float64) error {
	if rows == nil {
		return fmt.Errorf("parameter rows is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithDataErrorPolicyParameters(policy ErrorAction) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithFullLoadSubTasksParameters(count *float64) error {
	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithLobModeParameters(mode LobMode) error {
	if mode == "" {
		return fmt.Errorf("parameter mode is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithLoggingParameters(component *string, level LoggingLevel) error {
	if component == nil {
		return fmt.Errorf("parameter component is required, but nil was provided")
	}

	if level == "" {
		return fmt.Errorf("parameter level is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithLoggingEnabledParameters(enabled *bool) error {
	if enabled == nil {
		return fmt.Errorf("parameter enabled is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithRecoveryParameters(count *float64) error {
	if count == nil {
		return fmt.Errorf("parameter count is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskSettings) validateWithTargetTablePrepModeParameters(mode *string) error {
	if mode == nil {
		return fmt.Errorf("parameter mode is required, but nil was provided")
	}

	return nil
}

