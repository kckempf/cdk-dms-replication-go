//go:build !no_runtime_type_checking

package cdkdmsreplication

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_TableMappings) validateAddColumnParameters(schemaName *string, tableName *string, column *AddColumnDefinition) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if column == nil {
		return fmt.Errorf("parameter column is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(column, func() string { return "parameter column" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateAddPrefixToSchemaParameters(schemaName *string, prefix *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateAddPrefixToTableParameters(schemaName *string, tableName *string, prefix *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateAddSuffixToSchemaParameters(schemaName *string, suffix *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if suffix == nil {
		return fmt.Errorf("parameter suffix is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateAddSuffixToTableParameters(schemaName *string, tableName *string, suffix *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if suffix == nil {
		return fmt.Errorf("parameter suffix is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateExcludeSchemaParameters(schemaName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateExcludeTableParameters(schemaName *string, tableName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateExplicitTableParameters(schemaName *string, tableName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateIncludeSchemaParameters(schemaName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateIncludeTableParameters(schemaName *string, tableName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateMapToDynamoDbParameters(schemaName *string, tableName *string, options *DynamoDbObjectMappingOptions) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateRemoveColumnParameters(schemaName *string, tableName *string, columnName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if columnName == nil {
		return fmt.Errorf("parameter columnName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateRenameColumnParameters(schemaName *string, tableName *string, columnName *string, newName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if columnName == nil {
		return fmt.Errorf("parameter columnName is required, but nil was provided")
	}

	if newName == nil {
		return fmt.Errorf("parameter newName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateRenameSchemaParameters(schemaName *string, newName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if newName == nil {
		return fmt.Errorf("parameter newName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateRenameTableParameters(schemaName *string, tableName *string, newName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if newName == nil {
		return fmt.Errorf("parameter newName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateToLowerCaseColumnParameters(schemaName *string, tableName *string, columnName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if columnName == nil {
		return fmt.Errorf("parameter columnName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateToLowerCaseSchemaParameters(schemaName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateToLowerCaseTableParameters(schemaName *string, tableName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateToUpperCaseColumnParameters(schemaName *string, tableName *string, columnName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	if columnName == nil {
		return fmt.Errorf("parameter columnName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateToUpperCaseSchemaParameters(schemaName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TableMappings) validateToUpperCaseTableParameters(schemaName *string, tableName *string) error {
	if schemaName == nil {
		return fmt.Errorf("parameter schemaName is required, but nil was provided")
	}

	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

