package cdkdmsreplication


// Mapping from a source column (or DMS expression) to a DynamoDB partition or sort key attribute.
//
// Exactly one of `sourceColumn` or `value` must be set.
type DynamoDbKeyMapping struct {
	// DynamoDB attribute sub-type (`string`, `number`, or `binary`).
	AttributeSubType DynamoDbAttributeSubType `field:"required" json:"attributeSubType" yaml:"attributeSubType"`
	// Name of the partition/sort key attribute on the DynamoDB target item.
	TargetAttributeName *string `field:"required" json:"targetAttributeName" yaml:"targetAttributeName"`
	// Name of a single column on the relational source.
	//
	// The builder wraps this
	// in the DMS expression `${sourceColumn}`. Mutually exclusive with `value`.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// Raw DMS value expression, used verbatim.
	//
	// Use this for composite keys or
	// other expressions like `'CUSTOMER#${customer_id}'`. Mutually exclusive
	// with `sourceColumn`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

