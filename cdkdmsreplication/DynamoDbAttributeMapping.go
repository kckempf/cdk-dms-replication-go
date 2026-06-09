package cdkdmsreplication


// Mapping from a source column (or DMS expression) to a non-key attribute on the DynamoDB target.
//
// Exactly one of `sourceColumn` or `value` must be set.
// Source columns not listed in `attributeMappings` or `excludeColumns` pass
// through with the source column name.
type DynamoDbAttributeMapping struct {
	// DynamoDB attribute sub-type (`string`, `number`, or `binary`).
	AttributeSubType DynamoDbAttributeSubType `field:"required" json:"attributeSubType" yaml:"attributeSubType"`
	// Name of the attribute on the DynamoDB target item.
	TargetAttributeName *string `field:"required" json:"targetAttributeName" yaml:"targetAttributeName"`
	// Name of a single column on the relational source.
	//
	// The builder wraps this
	// in the DMS expression `${sourceColumn}`. Mutually exclusive with `value`.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// Raw DMS value expression, used verbatim.
	//
	// Use this for composite values or
	// other expressions like `'STATUS#${status}'`. Mutually exclusive with
	// `sourceColumn`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

