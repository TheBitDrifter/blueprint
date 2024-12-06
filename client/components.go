package blueprint_client

import "github.com/TheBitDrifter/warehouse"

type defaultComponents struct {
	SpriteLocations    spriteLocations
	ActiveSprite       warehouse.AccessibleComponent[ActiveSprite]
	Camera             warehouse.AccessibleComponent[Camera]
	RenderOptions      warehouse.AccessibleComponent[RenderOptions]
	ParallaxBackground warehouse.AccessibleComponent[ParallaxBackground]
}

type spriteLocations struct {
	Small warehouse.AccessibleComponent[SpriteLocationsSmall]
	Med   warehouse.AccessibleComponent[SpriteLocationsMed]
	Large warehouse.AccessibleComponent[SpriteLocationsLarge]
	XL    warehouse.AccessibleComponent[SpriteLocationsXL]
}

var Components = defaultComponents{
	ActiveSprite:       warehouse.FactoryNewComponent[ActiveSprite](),
	Camera:             warehouse.FactoryNewComponent[Camera](),
	RenderOptions:      warehouse.FactoryNewComponent[RenderOptions](),
	ParallaxBackground: warehouse.FactoryNewComponent[ParallaxBackground](),
	SpriteLocations: spriteLocations{
		Small: warehouse.FactoryNewComponent[SpriteLocationsSmall](),
		Med:   warehouse.FactoryNewComponent[SpriteLocationsMed](),
		Large: warehouse.FactoryNewComponent[SpriteLocationsLarge](),
		XL:    warehouse.FactoryNewComponent[SpriteLocationsXL](),
	},
}
