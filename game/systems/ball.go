package systems

import (
	"github.com/mrparano1d/ecs"
	ebitencomponents "github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/game/components"
)

func BallSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		query := ecs.NewQuery(ctx.World)
		for _, e := range query.Find(
			ecs.NewFilter(
				ecs.WithComponentFilter(
					&components.Ball{},
					&ebitencomponents.Position{},
					&ebitencomponents.Velocity{},
				),
			),
		) {
			position := ecs.GetComponent[*ebitencomponents.Position](ctx.World.Entities(), e)
			velocity := ecs.GetComponent[*ebitencomponents.Velocity](ctx.World.Entities(), e)

			position.X += velocity.X
			position.Y += velocity.Y
		}
	}
}
