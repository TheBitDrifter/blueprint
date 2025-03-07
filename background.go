package blueprint

import (
	blueprintclient "github.com/TheBitDrifter/blueprint/client"
	blueprintspatial "github.com/TheBitDrifter/blueprint/spatial"
	"github.com/TheBitDrifter/blueprint/vector"
	"github.com/TheBitDrifter/warehouse"
)

// ParallaxLayer defines a single layer in a parallax background
type ParallaxLayer struct {
	// SpritePath is the path to the sprite resource
	SpritePath string
	// SpeedX is the horizontal parallax speed multiplier
	SpeedX float64
	// SpeedY is the vertical parallax speed multiplier
	SpeedY float64
}

// ParallaxBackgroundBuilder provides a fluent API for creating parallax backgrounds
type ParallaxBackgroundBuilder struct {
	storage warehouse.Storage
	layers  []ParallaxLayer
	// Optional position offset for the entire background
	offset vector.Two
}

// NewParallaxBackgroundBuilder creates a new builder for parallax backgrounds
func NewParallaxBackgroundBuilder(sto warehouse.Storage) *ParallaxBackgroundBuilder {
	return &ParallaxBackgroundBuilder{
		storage: sto,
		layers:  []ParallaxLayer{},
	}
}

// WithOffset sets an optional position offset for the entire background
func (b *ParallaxBackgroundBuilder) WithOffset(offset vector.Two) *ParallaxBackgroundBuilder {
	b.offset = offset
	return b
}

// AddLayer adds a new layer to the parallax background
func (b *ParallaxBackgroundBuilder) AddLayer(spritePath string, speedX, speedY float64) *ParallaxBackgroundBuilder {
	b.layers = append(b.layers, ParallaxLayer{
		SpritePath: spritePath,
		SpeedX:     speedX,
		SpeedY:     speedY,
	})
	return b
}

// Build generates all layers and creates the parallax background
func (b *ParallaxBackgroundBuilder) Build() error {
	// Create the backgroundArchetype
	backgroundArchetype, err := b.storage.NewOrExistingArchetype(
		blueprintclient.Components.SpriteBundle,
		blueprintclient.Components.ParallaxBackground,
		blueprintspatial.Components.Position,
	)
	if err != nil {
		return err
	}

	// Handle empty layer list
	if len(b.layers) == 0 {
		return nil
	}

	// Generate each layer from the provided slice
	for _, layer := range b.layers {
		sprite := blueprintclient.NewSpriteBundle().AddSprite(layer.SpritePath, true)

		// Apply offset if specified
		if b.offset.X != 0 || b.offset.Y != 0 {
			// Backgrounds only use first index
			sprite.Blueprints[0].Config.Offset = b.offset
		}

		err = backgroundArchetype.Generate(
			1,
			sprite,
			blueprintclient.ParallaxBackground{
				SpeedX: layer.SpeedX,
				SpeedY: layer.SpeedY,
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateStillBackground is a utility function for creating a non-parallax (static) background
func CreateStillBackground(sto warehouse.Storage, spritePath string) error {
	backgroundArchetype, err := sto.NewOrExistingArchetype(
		blueprintclient.Components.SpriteBundle,
		blueprintclient.Components.ParallaxBackground,
		blueprintspatial.Components.Position,
	)
	if err != nil {
		return err
	}

	return backgroundArchetype.Generate(
		1,
		blueprintclient.NewSpriteBundle().AddSprite(spritePath, true),
		blueprintclient.ParallaxBackground{
			SpeedX: 0,
			SpeedY: 0,
		},
	)
}
