package cdkdmsreplication


// LOB (Large Object) handling mode.
type LobMode string

const (
	// No LOB support.
	LobMode_NONE LobMode = "NONE"
	// All LOBs are migrated as inline data (full LOB mode).
	LobMode_FULL_LOB LobMode = "FULL_LOB"
	// LOBs are truncated at a configurable size.
	LobMode_LIMITED_LOB LobMode = "LIMITED_LOB"
)

