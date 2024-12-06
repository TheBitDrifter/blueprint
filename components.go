package blueprint

import "github.com/TheBitDrifter/warehouse"

type Position struct {
	X, Y float64
}

// ----------------------------------------------------

type defaultComponents struct {
	Position warehouse.AccessibleComponent[Position]
}

var Components = defaultComponents{
	Position: warehouse.FactoryNewComponent[Position](),
}
