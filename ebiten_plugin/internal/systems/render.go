package systems

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"
)

func RenderSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		surface := ecs.GetResource[interfaces.Surface](ctx.Resources)
		query := ecs.NewQuery(ctx.World)
		for _, e := range query.Find(ecs.NewFilter(ecs.WithComponentFilter(&components.Position{}, &components.Asset{}))) {
			position := ecs.GetComponent[*components.Position](ctx.World.Entities(), e)
			asset := ecs.GetComponent[*components.Asset](ctx.World.Entities(), e)

			asset.Handle.Translate(position.X, position.Y)
			asset.Handle.Draw(surface)
		}
	}
}
