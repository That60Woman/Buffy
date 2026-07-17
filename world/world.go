package world

// World is Buffy's internal model of reality.
//
// Reality exists independently.
//
// The World is Buffy's continuously evolving,
// high-fidelity model of that reality.
//
// Everything else is derived from it.
type World struct {
}

// New creates an empty World.
func New() *World {
	return &World{}
}
