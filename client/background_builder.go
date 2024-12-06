package blueprint_client

import (
	"log"

	"github.com/TheBitDrifter/blueprint"
	"github.com/TheBitDrifter/warehouse"
)

var BackgroundBuilder = backgroundBuilder{}

type backgroundBuilder struct{}

type parallaxEntities []warehouse.Entity

func (backgroundBuilder) AddBackgrounds(sto warehouse.Storage, paths ...string) error {
	entities, err := sto.NewEntities(
		len(paths),
		Components.SpriteLocations.Small, Components.ActiveSprite, blueprint.Components.Position,
	)
	if err != nil {
		return err
	}
	for i, en := range entities {
		imageLocations := Components.SpriteLocations.Small.GetFromEntity(en)
		imageLocations[0].Key = paths[i]
	}
	return nil
}

func (backgroundBuilder) AddParallaxBackgrounds(sto warehouse.Storage, paths ...string) (parallaxEntities, error) {
	entities, err := sto.NewEntities(
		len(paths),
		Components.SpriteLocations.Small,
		Components.ActiveSprite,
		Components.ParallaxBackground,
		blueprint.Components.Position,
	)
	if err != nil {
		return parallaxEntities(entities), err
	}
	for i, en := range entities {
		imageLocations := Components.SpriteLocations.Small.GetFromEntity(en)
		imageLocations[0].Key = paths[i]

	}
	return parallaxEntities(entities), nil
}

func (pe parallaxEntities) Set(backgroundSettings ...ParallaxBackground) error {
	if len(pe) != len(backgroundSettings) {
		log.Fatal("todo")
	}
	for i, backgroundSetting := range backgroundSettings {
		existingSetting := Components.ParallaxBackground.GetFromEntity(pe[i])
		existingSetting.SpeedX = backgroundSetting.SpeedX
		existingSetting.SpeedY = backgroundSetting.SpeedY
		existingSetting.LoopX = backgroundSetting.LoopX
		existingSetting.LoopY = backgroundSetting.LoopY
	}
	return nil
}
