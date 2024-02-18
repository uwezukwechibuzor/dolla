package types

const (
	// ModuleName defines the module name
	ModuleName = "dollar"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dollar"
)

var (
	ParamsKey = []byte("p_dollar")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PostKey      = "Post/value/"
	PostCountKey = "Post/count/"
)
