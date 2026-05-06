package cdkdmsreplication


// Represents a single, fully built rule in the table-mappings JSON.
type TableMappingRule struct {
	ObjectLocator *RuleObjectLocatorValue `field:"required" json:"objectLocator" yaml:"objectLocator"`
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	OldValue *string `field:"optional" json:"oldValue" yaml:"oldValue"`
	Value *string `field:"optional" json:"value" yaml:"value"`
}

