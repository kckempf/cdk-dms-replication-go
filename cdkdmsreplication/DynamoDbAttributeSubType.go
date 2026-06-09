package cdkdmsreplication


// DynamoDB attribute sub-type for object-mapping rules.
type DynamoDbAttributeSubType string

const (
	DynamoDbAttributeSubType_STRING DynamoDbAttributeSubType = "STRING"
	DynamoDbAttributeSubType_NUMBER DynamoDbAttributeSubType = "NUMBER"
	DynamoDbAttributeSubType_BINARY DynamoDbAttributeSubType = "BINARY"
)

