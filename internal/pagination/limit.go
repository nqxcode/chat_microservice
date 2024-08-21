package pagination

const DefaultLimit = uint64(30)

// Limit limit
type Limit struct {
	Offset uint64
	Limit  uint64
}
