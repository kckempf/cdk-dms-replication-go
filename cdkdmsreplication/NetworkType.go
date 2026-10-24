package cdkdmsreplication


// IP addressing mode for a replication instance.
type NetworkType string

const (
	// IPv4 only.
	NetworkType_IPV4 NetworkType = "IPV4"
	// Dual-stack: IPv4 and IPv6.
	//
	// Requires subnets with IPv6 CIDR blocks.
	NetworkType_DUAL NetworkType = "DUAL"
)

