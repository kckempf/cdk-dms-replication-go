package cdkdmsreplication


// Oracle number datatype scale option.
type OracleNumberDatatypeScale string

const (
	// Use -1 to preserve the original Oracle precision.
	OracleNumberDatatypeScale_FLOAT OracleNumberDatatypeScale = "FLOAT"
	// Use -2 to map NUMBER to DOUBLE.
	OracleNumberDatatypeScale_DOUBLE OracleNumberDatatypeScale = "DOUBLE"
)

