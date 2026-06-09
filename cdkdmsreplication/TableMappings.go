package cdkdmsreplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kckempf/cdk-dms-replication-go/cdkdmsreplication/jsii"
)

// Fluent builder for DMS table mappings.
//
// Example:
//   const mappings = new TableMappings()
//     .includeSchema('public')
//     .includeTable('public', 'orders')
//     .excludeTable('public', 'audit_log')
//     .renameSchema('public', 'prod')
//     .toLowerCaseTable('public', '%')
//     .toJson();
//
type TableMappings interface {
	// Add a new column to a table.
	//
	// Example:
	//   mappings.addColumn('public', 'orders', {
	//     columnName: 'migration_ts',
	//     columnType: ColumnDataType.DATETIME,
	//     expression: '$timestamp',
	//   });
	//
	AddColumn(schemaName *string, tableName *string, column *AddColumnDefinition) TableMappings
	// Add a prefix to schema names.
	AddPrefixToSchema(schemaName *string, prefix *string) TableMappings
	// Add a prefix to table names.
	AddPrefixToTable(schemaName *string, tableName *string, prefix *string) TableMappings
	// Add a suffix to schema names.
	AddSuffixToSchema(schemaName *string, suffix *string) TableMappings
	// Add a suffix to table names.
	AddSuffixToTable(schemaName *string, tableName *string, suffix *string) TableMappings
	// Exclude all tables in a schema.
	ExcludeSchema(schemaName *string) TableMappings
	// Exclude a specific table (or wildcard) within a schema.
	ExcludeTable(schemaName *string, tableName *string) TableMappings
	// Explicitly include a single table.
	//
	// Unlike `include`, `explicit` means DMS
	// only migrates this one table regardless of other `include` rules.
	ExplicitTable(schemaName *string, tableName *string) TableMappings
	// Include all tables in a schema.
	//
	// Use `%` as a wildcard for `schemaName` to include all schemas.
	IncludeSchema(schemaName *string) TableMappings
	// Include a specific table (or a wildcard pattern) within a schema.
	//
	// Use `%` for `tableName` to match all tables in the schema.
	IncludeTable(schemaName *string, tableName *string) TableMappings
	// Map a relational source table to a DynamoDB target table.
	//
	// Emits a DMS `object-mapping` rule with `rule-action: map-record-to-record`.
	// DMS requires this rule type when the target endpoint is DynamoDB; the
	// partition key (and optional sort key) tell DMS how to build the item key
	// from source columns.
	//
	// Source columns not listed in `attributeMappings` or `excludeColumns` are
	// migrated with the source column name as the attribute name.
	//
	// For each key or attribute mapping, set either `sourceColumn` (a single
	// column wrapped as `${col}`) or `value` (a raw DMS expression, e.g.
	// `'CUSTOMER#${customer_id}'` for composite keys). Exactly one is required.
	//
	// Calling this method twice with the same `schemaName`/`tableName` emits two
	// separate rules; DMS will reject duplicate object-mapping rules at deploy
	// time. Call it once per source table.
	//
	// Example:
	//   const mappings = new TableMappings()
	//     .includeTable('public', 'orders')
	//     .mapToDynamoDb('public', 'orders', {
	//       targetTableName: 'Orders',
	//       // Composite partition key from a literal prefix + source column:
	//       partitionKey: {
	//         value: 'CUSTOMER#${customer_id}',
	//         targetAttributeName: 'PK',
	//         attributeSubType: DynamoDbAttributeSubType.STRING,
	//       },
	//       // Bare source column for the sort key:
	//       sortKey: {
	//         sourceColumn: 'created_at',
	//         targetAttributeName: 'CreatedAt',
	//         attributeSubType: DynamoDbAttributeSubType.STRING,
	//       },
	//       excludeColumns: ['internal_flag'],
	//       attributeMappings: [
	//         {
	//           sourceColumn: 'customer_id',
	//           targetAttributeName: 'CustomerId',
	//           attributeSubType: DynamoDbAttributeSubType.STRING,
	//         },
	//       ],
	//     })
	//     .toJson();
	//
	MapToDynamoDb(schemaName *string, tableName *string, options *DynamoDbObjectMappingOptions) TableMappings
	// Remove a column from a table.
	RemoveColumn(schemaName *string, tableName *string, columnName *string) TableMappings
	// Rename a column in a table.
	RenameColumn(schemaName *string, tableName *string, columnName *string, newName *string) TableMappings
	// Rename a schema.
	RenameSchema(schemaName *string, newName *string) TableMappings
	// Rename a table.
	RenameTable(schemaName *string, tableName *string, newName *string) TableMappings
	// Serialise the accumulated rules to the JSON string expected by DMS.
	//
	// Passes the result directly to `replicationTaskSettings` or
	// `DmsReplicationTask.tableMappings`.
	ToJson() *string
	// Convert matching column names to lowercase.
	ToLowerCaseColumn(schemaName *string, tableName *string, columnName *string) TableMappings
	// Convert all schema names to lowercase.
	ToLowerCaseSchema(schemaName *string) TableMappings
	// Convert matching table names to lowercase.
	//
	// Use `%` to match all tables.
	ToLowerCaseTable(schemaName *string, tableName *string) TableMappings
	// Convert matching column names to uppercase.
	ToUpperCaseColumn(schemaName *string, tableName *string, columnName *string) TableMappings
	// Convert all schema names to uppercase.
	ToUpperCaseSchema(schemaName *string) TableMappings
	// Convert matching table names to uppercase.
	ToUpperCaseTable(schemaName *string, tableName *string) TableMappings
}

// The jsii proxy struct for TableMappings
type jsiiProxy_TableMappings struct {
	_ byte // padding
}

func NewTableMappings() TableMappings {
	_init_.Initialize()

	j := jsiiProxy_TableMappings{}

	_jsii_.Create(
		"cdk-dms-replication.TableMappings",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewTableMappings_Override(t TableMappings) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dms-replication.TableMappings",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TableMappings) AddColumn(schemaName *string, tableName *string, column *AddColumnDefinition) TableMappings {
	if err := t.validateAddColumnParameters(schemaName, tableName, column); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"addColumn",
		[]interface{}{schemaName, tableName, column},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) AddPrefixToSchema(schemaName *string, prefix *string) TableMappings {
	if err := t.validateAddPrefixToSchemaParameters(schemaName, prefix); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"addPrefixToSchema",
		[]interface{}{schemaName, prefix},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) AddPrefixToTable(schemaName *string, tableName *string, prefix *string) TableMappings {
	if err := t.validateAddPrefixToTableParameters(schemaName, tableName, prefix); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"addPrefixToTable",
		[]interface{}{schemaName, tableName, prefix},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) AddSuffixToSchema(schemaName *string, suffix *string) TableMappings {
	if err := t.validateAddSuffixToSchemaParameters(schemaName, suffix); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"addSuffixToSchema",
		[]interface{}{schemaName, suffix},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) AddSuffixToTable(schemaName *string, tableName *string, suffix *string) TableMappings {
	if err := t.validateAddSuffixToTableParameters(schemaName, tableName, suffix); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"addSuffixToTable",
		[]interface{}{schemaName, tableName, suffix},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ExcludeSchema(schemaName *string) TableMappings {
	if err := t.validateExcludeSchemaParameters(schemaName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"excludeSchema",
		[]interface{}{schemaName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ExcludeTable(schemaName *string, tableName *string) TableMappings {
	if err := t.validateExcludeTableParameters(schemaName, tableName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"excludeTable",
		[]interface{}{schemaName, tableName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ExplicitTable(schemaName *string, tableName *string) TableMappings {
	if err := t.validateExplicitTableParameters(schemaName, tableName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"explicitTable",
		[]interface{}{schemaName, tableName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) IncludeSchema(schemaName *string) TableMappings {
	if err := t.validateIncludeSchemaParameters(schemaName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"includeSchema",
		[]interface{}{schemaName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) IncludeTable(schemaName *string, tableName *string) TableMappings {
	if err := t.validateIncludeTableParameters(schemaName, tableName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"includeTable",
		[]interface{}{schemaName, tableName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) MapToDynamoDb(schemaName *string, tableName *string, options *DynamoDbObjectMappingOptions) TableMappings {
	if err := t.validateMapToDynamoDbParameters(schemaName, tableName, options); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"mapToDynamoDb",
		[]interface{}{schemaName, tableName, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) RemoveColumn(schemaName *string, tableName *string, columnName *string) TableMappings {
	if err := t.validateRemoveColumnParameters(schemaName, tableName, columnName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"removeColumn",
		[]interface{}{schemaName, tableName, columnName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) RenameColumn(schemaName *string, tableName *string, columnName *string, newName *string) TableMappings {
	if err := t.validateRenameColumnParameters(schemaName, tableName, columnName, newName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"renameColumn",
		[]interface{}{schemaName, tableName, columnName, newName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) RenameSchema(schemaName *string, newName *string) TableMappings {
	if err := t.validateRenameSchemaParameters(schemaName, newName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"renameSchema",
		[]interface{}{schemaName, newName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) RenameTable(schemaName *string, tableName *string, newName *string) TableMappings {
	if err := t.validateRenameTableParameters(schemaName, tableName, newName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"renameTable",
		[]interface{}{schemaName, tableName, newName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ToJson() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ToLowerCaseColumn(schemaName *string, tableName *string, columnName *string) TableMappings {
	if err := t.validateToLowerCaseColumnParameters(schemaName, tableName, columnName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"toLowerCaseColumn",
		[]interface{}{schemaName, tableName, columnName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ToLowerCaseSchema(schemaName *string) TableMappings {
	if err := t.validateToLowerCaseSchemaParameters(schemaName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"toLowerCaseSchema",
		[]interface{}{schemaName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ToLowerCaseTable(schemaName *string, tableName *string) TableMappings {
	if err := t.validateToLowerCaseTableParameters(schemaName, tableName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"toLowerCaseTable",
		[]interface{}{schemaName, tableName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ToUpperCaseColumn(schemaName *string, tableName *string, columnName *string) TableMappings {
	if err := t.validateToUpperCaseColumnParameters(schemaName, tableName, columnName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"toUpperCaseColumn",
		[]interface{}{schemaName, tableName, columnName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ToUpperCaseSchema(schemaName *string) TableMappings {
	if err := t.validateToUpperCaseSchemaParameters(schemaName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"toUpperCaseSchema",
		[]interface{}{schemaName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableMappings) ToUpperCaseTable(schemaName *string, tableName *string) TableMappings {
	if err := t.validateToUpperCaseTableParameters(schemaName, tableName); err != nil {
		panic(err)
	}
	var returns TableMappings

	_jsii_.Invoke(
		t,
		"toUpperCaseTable",
		[]interface{}{schemaName, tableName},
		&returns,
	)

	return returns
}

