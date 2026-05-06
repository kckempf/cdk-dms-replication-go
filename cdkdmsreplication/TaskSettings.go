package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/jsii"
)

// Fluent builder for DMS task settings.
//
// Example:
//   const settings = new TaskSettings()
//     .withLobMode(LobMode.LIMITED_LOB, 32)
//     .withFullLoadSubTasks(16)
//     .withBatchApply(true, 5, 60)
//     .withDataErrorPolicy(ErrorAction.IGNORE_RECORD, 100)
//     .withLogging('SOURCE_UNLOAD', LoggingLevel.LOGGER_SEVERITY_DEFAULT)
//     .toJson();
//
type TaskSettings interface {
	// Produce the JSON string that DMS expects for `replicationTaskSettings`.
	ToJson() *string
	// Enable or disable CDC batch apply mode and configure its timeouts.
	WithBatchApply(enabled *bool, minSeconds *float64, maxSeconds *float64) TaskSettings
	// Set the commit rate (number of rows per commit) during full load.
	WithCommitRate(rows *float64) TaskSettings
	// Configure the action taken on data errors.
	WithDataErrorPolicy(policy ErrorAction, escalationCount *float64, escalationPolicy ErrorAction) TaskSettings
	// Set the maximum number of tables loaded in parallel.
	WithFullLoadSubTasks(count *float64) TaskSettings
	// Configure LOB handling.
	WithLobMode(mode LobMode, maxSizeKb *float64) TaskSettings
	// Set the logging level for a named DMS log component.
	//
	// Common component names: SOURCE_UNLOAD, TARGET_LOAD, TASK_MANAGER,
	// SOURCE_CAPTURE, TARGET_APPLY, REST_SERVER, TRANSFORMATION.
	WithLogging(component *string, level LoggingLevel) TaskSettings
	// Enable or disable CloudWatch logging for the task.
	// Default: true.
	//
	WithLoggingEnabled(enabled *bool) TaskSettings
	// Configure recovery from transient errors.
	WithRecovery(count *float64, interval *float64) TaskSettings
	// Set the full-load target table prepare mode.
	//
	// Common values: "DROP_AND_CREATE", "TRUNCATE_BEFORE_LOAD", "DO_NOTHING".
	WithTargetTablePrepMode(mode *string) TaskSettings
}

// The jsii proxy struct for TaskSettings
type jsiiProxy_TaskSettings struct {
	_ byte // padding
}

func NewTaskSettings() TaskSettings {
	_init_.Initialize()

	j := jsiiProxy_TaskSettings{}

	_jsii_.Create(
		"cdk-dms-replication.TaskSettings",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewTaskSettings_Override(t TaskSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dms-replication.TaskSettings",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TaskSettings) ToJson() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithBatchApply(enabled *bool, minSeconds *float64, maxSeconds *float64) TaskSettings {
	if err := t.validateWithBatchApplyParameters(enabled); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withBatchApply",
		[]interface{}{enabled, minSeconds, maxSeconds},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithCommitRate(rows *float64) TaskSettings {
	if err := t.validateWithCommitRateParameters(rows); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withCommitRate",
		[]interface{}{rows},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithDataErrorPolicy(policy ErrorAction, escalationCount *float64, escalationPolicy ErrorAction) TaskSettings {
	if err := t.validateWithDataErrorPolicyParameters(policy); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withDataErrorPolicy",
		[]interface{}{policy, escalationCount, escalationPolicy},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithFullLoadSubTasks(count *float64) TaskSettings {
	if err := t.validateWithFullLoadSubTasksParameters(count); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withFullLoadSubTasks",
		[]interface{}{count},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithLobMode(mode LobMode, maxSizeKb *float64) TaskSettings {
	if err := t.validateWithLobModeParameters(mode); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withLobMode",
		[]interface{}{mode, maxSizeKb},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithLogging(component *string, level LoggingLevel) TaskSettings {
	if err := t.validateWithLoggingParameters(component, level); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withLogging",
		[]interface{}{component, level},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithLoggingEnabled(enabled *bool) TaskSettings {
	if err := t.validateWithLoggingEnabledParameters(enabled); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withLoggingEnabled",
		[]interface{}{enabled},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithRecovery(count *float64, interval *float64) TaskSettings {
	if err := t.validateWithRecoveryParameters(count); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withRecovery",
		[]interface{}{count, interval},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskSettings) WithTargetTablePrepMode(mode *string) TaskSettings {
	if err := t.validateWithTargetTablePrepModeParameters(mode); err != nil {
		panic(err)
	}
	var returns TaskSettings

	_jsii_.Invoke(
		t,
		"withTargetTablePrepMode",
		[]interface{}{mode},
		&returns,
	)

	return returns
}

