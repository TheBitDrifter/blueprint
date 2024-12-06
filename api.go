package blueprint

import "github.com/TheBitDrifter/warehouse"

type (
	Plan = func(warehouse.Storage) error
)
