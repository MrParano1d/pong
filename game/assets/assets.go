package assets

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/components"
	"github.com/mrparano1d/pong/opengl/shapes"
	"github.com/mrparano1d/pong/opengl/time"
	"github.com/mrparano1d/pong/opengl/window"
	"image/color"
)

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.Spawn().Insert(
			&components.Asset{
				Handle: shapes.NewRectangle(0, 0, 200, 200,
					color.RGBA{
						R: 0,
						G: 255,
						B: 0,
						A: 255,
					},
				),
			},
		).Insert(&components.Position{
			X: 0,
			Y: 0,
		}).Insert(&components.Velocity{
			X: 150,
			Y: 150,
		})

		commands.Spawn().Insert(
			&components.Asset{
				Handle: shapes.NewRectangle(0, 0, 200, 200,
					color.RGBA{
						R: 255,
						G: 255,
						B: 0,
						A: 255,
					},
				),
			},
		).Insert(&components.Position{
			X: 300,
			Y: 150,
		}).Insert(&components.Velocity{
			X: 150,
			Y: 150,
		})
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {
		win := ecs.GetResource[*window.Resource](ctx.Resources)
		t := ecs.GetResource[*time.Resource](ctx.Resources)

		q := ecs.NewQuery(ctx.World)
		for _, eID := range q.Find(ecs.NewFilter(ecs.WithComponentFilter(&components.Asset{}, &components.Position{}, &components.Velocity{}))) {
			position := ecs.GetComponent[*components.Position](ctx.World.Entities(), eID)
			asset := ecs.GetComponent[*components.Asset](ctx.World.Entities(), eID)
			velocity := ecs.GetComponent[*components.Velocity](ctx.World.Entities(), eID)
			position.X += velocity.X * float32(t.Delta())
			position.Y += velocity.Y * float32(t.Delta())

			if position.X+asset.Handle.Width() >= win.Width || position.X <= 0 {
				velocity.X *= -1
			}
			if position.Y+asset.Handle.Height() >= win.Height || position.Y <= 0 {
				velocity.Y *= -1
			}
		}
	}
}
