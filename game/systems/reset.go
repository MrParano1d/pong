package systems

import (
	"github.com/mrparano1d/ecs"
	ebitencomponents "github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/game/assets"
	"github.com/mrparano1d/pong/game/components"
	gameevents "github.com/mrparano1d/pong/game/events"
)

func ResetSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		reader := ctx.EventReader(gameevents.ScoredEvent{})
		for reader.Next() {
			_ = reader.Read()

			playerQuery := ecs.NewQuery(ctx.World)
			for _, p := range playerQuery.Find(
				ecs.NewFilter(
					ecs.WithComponentFilter(
						&components.Player{}, &ebitencomponents.Position{},
					),
				),
			) {
				position := ecs.GetComponent[*ebitencomponents.Position](ctx.World.Entities(), p)

				position.Y = assets.PlayerPositionY
			}

			ballQuery := ecs.NewQuery(ctx.World)
			for _, b := range ballQuery.Find(
				ecs.NewFilter(
					ecs.WithComponentFilter(
						&components.Ball{},
						&ebitencomponents.Position{},
					),
				),
			) {
				position := ecs.GetComponent[*ebitencomponents.Position](ctx.World.Entities(), b)

				position.X = assets.BallPositionX
				position.Y = assets.BallPositionY
			}
		}
	}
}
