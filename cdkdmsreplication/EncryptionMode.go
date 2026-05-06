package cdkdmsreplication


// Encryption mode used when writing data to an S3 bucket or Redshift cluster.
type EncryptionMode string

const (
	// Server-side encryption using S3-managed keys (SSE-S3).
	EncryptionMode_SSE_S3 EncryptionMode = "SSE_S3"
	// Server-side encryption using AWS KMS-managed keys (SSE-KMS).
	EncryptionMode_SSE_KMS EncryptionMode = "SSE_KMS"
)

