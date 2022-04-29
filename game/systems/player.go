package systems

import (
	"github.com/mrparano1d/ecs"
	ebitencomponents "github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/ebiten_plugin/input"
	"github.com/mrparano1d/pong/game/components"
)

func PlayerSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		inputManager := ecs.GetResource[*input.InputManager](ctx.Resources)
		query := ecs.NewQuery(ctx.World)
		for _, e := range query.Find(
			ecs.NewFilter(
				ecs.WithComponentFilter(
					&components.Player{},
					&ebitencomponents.Position{},
					&ebitencomponents.Velocity{},
				),
			),
		) {
			player := ecs.GetComponent[*components.Player](ctx.World.Entities(), e)
			position := ecs.GetComponent[*ebitencomponents.Position](ctx.World.Entities(), e)
			velocity := ecs.GetComponent[*ebitencomponents.Velocity](ctx.World.Entities(), e)

			if inputManager.GetInputService().IsKeyPressed(player.UpKey) {
				position.Y -= velocity.Y
			} else if inputManager.GetInputService().IsKeyPressed(player.DownKey) {
				position.Y += velocity.Y
			}
		}
	}
}
