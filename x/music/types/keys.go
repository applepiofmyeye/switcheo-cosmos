package types

const (
	// ModuleName defines the module name
	ModuleName = "music"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_music"
)

var (
	ParamsKey = []byte("p_music")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
