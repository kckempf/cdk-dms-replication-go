package cdkdmsreplication


// Encoding type for S3 parquet output.
type EncodingType string

const (
	EncodingType_PLAIN EncodingType = "PLAIN"
	EncodingType_PLAIN_DICTIONARY EncodingType = "PLAIN_DICTIONARY"
	EncodingType_RLE_DICTIONARY EncodingType = "RLE_DICTIONARY"
)

