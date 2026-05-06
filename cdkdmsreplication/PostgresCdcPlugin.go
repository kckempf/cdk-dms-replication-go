package cdkdmsreplication


// PostgreSQL CDC plugin.
type PostgresCdcPlugin string

const (
	PostgresCdcPlugin_PG_LOGICAL PostgresCdcPlugin = "PG_LOGICAL"
	PostgresCdcPlugin_TEST_DECODING PostgresCdcPlugin = "TEST_DECODING"
)

