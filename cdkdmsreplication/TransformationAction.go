package cdkdmsreplication


// Table-level transformation actions.
type TransformationAction string

const (
	// Convert the name to lowercase.
	TransformationAction_CONVERT_LOWERCASE TransformationAction = "CONVERT_LOWERCASE"
	// Convert the name to uppercase.
	TransformationAction_CONVERT_UPPERCASE TransformationAction = "CONVERT_UPPERCASE"
	// Add a prefix to the name.
	TransformationAction_ADD_PREFIX TransformationAction = "ADD_PREFIX"
	// Remove a prefix from the name.
	TransformationAction_REMOVE_PREFIX TransformationAction = "REMOVE_PREFIX"
	// Add a suffix to the name.
	TransformationAction_ADD_SUFFIX TransformationAction = "ADD_SUFFIX"
	// Remove a suffix from the name.
	TransformationAction_REMOVE_SUFFIX TransformationAction = "REMOVE_SUFFIX"
	// Rename the object.
	TransformationAction_RENAME TransformationAction = "RENAME"
	// Remove the column.
	TransformationAction_REMOVE_COLUMN TransformationAction = "REMOVE_COLUMN"
	// Add a column.
	TransformationAction_ADD_COLUMN TransformationAction = "ADD_COLUMN"
	// Include the column.
	TransformationAction_INCLUDE_COLUMN TransformationAction = "INCLUDE_COLUMN"
)

