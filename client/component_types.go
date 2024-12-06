package blueprint_client

import (
	"github.com/TheBitDrifter/blueprint"
	"github.com/TheBitDrifter/warehouse"
)

type (
	SpriteLocationsSmall [5]warehouse.CacheLocation
	SpriteLocationsMed   [10]warehouse.CacheLocation
	SpriteLocationsLarge [20]warehouse.CacheLocation
	SpriteLocationsXL    [40]warehouse.CacheLocation
)

type ActiveSprite struct {
	Index int
}

type RenderOptions struct {
	TranslateX, TranslateY float64
}

type ParallaxBackground struct {
	SpeedX, SpeedY float64
	LoopX, LoopY   bool
}

type Camera struct {
	Height, Width int
	Zoom          float64
	Positions     struct {
		Screen blueprint.Position
		Local  blueprint.Position
	}
}
