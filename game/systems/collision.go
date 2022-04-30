package systems

import (
	"github.com/mrparano1d/ecs"
	ebitencomponents "github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/ebiten_plugin/events"
	"github.com/mrparano1d/pong/game/components"
)

func CollisionSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		reader := ctx.EventReader(events.CollisionEvent{})
		for reader.Next() {
			ev := reader.Read().(events.CollisionEvent)

			e1ball := ecs.GetComponent[*components.Ball](ctx.World.Entities(), ev.Entity1)
			e1player := ecs.GetComponent[*components.Player](ctx.World.Entities(), ev.Entity1)
			e1position := ecs.GetComponent[*ebitencomponents.Position](ctx.World.Entities(), ev.Entity1)
			e1collision := ecs.GetComponent[*ebitencomponents.Collision](ctx.World.Entities(), ev.Entity1)

			e2ball := ecs.GetComponent[*components.Ball](ctx.World.Entities(), ev.Entity2)
			e2player := ecs.GetComponent[*components.Player](ctx.World.Entities(), ev.Entity2)
			e2position := ecs.GetComponent[*ebitencomponents.Position](ctx.World.Entities(), ev.Entity2)
			e2collision := ecs.GetComponent[*ebitencomponents.Collision](ctx.World.Entities(), ev.Entity2)

			if e1ball != nil && e2player != nil {
				// entity 1 is the ball
				ballVelocity := ecs.GetComponent[*ebitencomponents.Velocity](ctx.World.Entities(), ev.Entity1)
				ballVelocity.X *= -1

				if ev.Side == events.CollisionRight {
					e1position.X = e2position.X - e1collision.Width
				} else if ev.Side == events.CollisionLeft {
					e1position.X = e2position.X + e2collision.Width
				}

			} else if e1player != nil && e2ball != nil {
				// entity 2 is the ball
				ballVelocity := ecs.GetComponent[*ebitencomponents.Velocity](ctx.World.Entities(), ev.Entity2)
				ballVelocity.X *= -1

				if ev.Side == events.CollisionLeft {
					e2position.X = e1position.X - e2collision.Width
				} else if ev.Side == events.CollisionRight {
					e2position.X = e1position.X + e1collision.Width
				}
			}
		}
	}
}
