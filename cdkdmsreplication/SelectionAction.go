package cdkdmsreplication


// Whether the rule selects or excludes objects.
type SelectionAction string

const (
	SelectionAction_INCLUDE SelectionAction = "INCLUDE"
	SelectionAction_EXCLUDE SelectionAction = "EXCLUDE"
	SelectionAction_EXPLICIT SelectionAction = "EXPLICIT"
)

