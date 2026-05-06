package cdkdmsreplication


// The object type a rule targets.
type RuleObjectLocator string

const (
	RuleObjectLocator_SCHEMA RuleObjectLocator = "SCHEMA"
	RuleObjectLocator_TABLE RuleObjectLocator = "TABLE"
	RuleObjectLocator_COLUMN RuleObjectLocator = "COLUMN"
	RuleObjectLocator_TABLE_TABLESPACE RuleObjectLocator = "TABLE_TABLESPACE"
	RuleObjectLocator_INDEX_TABLESPACE RuleObjectLocator = "INDEX_TABLESPACE"
)

