package systems

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/assets"
	"github.com/mrparano1d/pong/ebiten_plugin/input"
	"github.com/mrparano1d/pong/ebiten_plugin/input/ebiten"
	"github.com/mrparano1d/pong/ebiten_plugin/resources"
)

func PrepareSetup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(
			func(resourceMap ecs.ResourceMap) {
				ecs.AddResource[*assets.Manager](resourceMap, assets.NewManager())
				ecs.AddResource[*input.InputManager](resourceMap, input.NewInputManager(ebiten.InputService{}))
				ecs.AddResource[*resources.WorldBoundaries](
					resourceMap, &resources.WorldBoundaries{
						Width:  512,
						Height: 256,
					},
				)
			},
		)
	}
}
