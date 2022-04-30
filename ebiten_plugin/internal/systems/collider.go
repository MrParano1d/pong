package systems

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/ebiten_plugin/events"
	"github.com/mrparano1d/pong/ebiten_plugin/resources"
	"image"
)

func ColliderSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		query := ecs.NewQuery(ctx.World)
		worldBoundaries := ecs.GetResource[*resources.WorldBoundaries](ctx.Resources)

		collisions := map[ecs.Entity]ecs.Entity{}

		for e1, _ := range ctx.World.Entities() {
			e1position := ecs.GetComponent[*components.Position](ctx.World.Entities(), e1)
			if e1position == nil {
				continue
			}
			e1collision := ecs.GetComponent[*components.Collision](ctx.World.Entities(), e1)
			if e1collision == nil {
				continue
			}
			for e2, _ := range ctx.World.Entities() {
				// don't check for collisions with yourself
				if e1 == e2 {
					continue
				}
				// skip if this collision has been already detected
				if e, ok := collisions[e2]; ok && e == e1 {
					continue
				}
				e2position := ecs.GetComponent[*components.Position](ctx.World.Entities(), e2)
				if e2position == nil {
					continue
				}
				e2collision := ecs.GetComponent[*components.Collision](ctx.World.Entities(), e2)
				if e2collision == nil {
					continue
				}

				r1 := image.Rect(
					int(e1position.X),
					int(e1position.Y),
					int(e1position.X+e1collision.Width),
					int(e1position.Y+e1collision.Height),
				)
				r2 := image.Rect(
					int(e2position.X),
					int(e2position.Y),
					int(e2position.X+e2collision.Width),
					int(e2position.Y+e2collision.Height),
				)

				if r1.Overlaps(r2) {
					collisions[e1] = e2

					side := events.CollisionTop
					if r1.Min.X > r2.Min.X {
						side = events.CollisionLeft
					} else {
						side = events.CollisionRight
					}

					ctx.EventWriter(
						events.CollisionEvent{},
					).SendOnce(
						events.CollisionEvent{
							Entity1: e1,
							Entity2: e2,
							Side:    side,
						},
					)
				}
			}
		}

		for _, e := range query.Find(
			ecs.NewFilter(
				ecs.WithComponentFilter(
					&components.WorldCollider{},
					&components.Position{},
					&components.Velocity{},
					&components.Asset{},
				),
			),
		) {
			asset := ecs.GetComponent[*components.Asset](ctx.World.Entities(), e)
			position := ecs.GetComponent[*components.Position](ctx.World.Entities(), e)
			velocity := ecs.GetComponent[*components.Velocity](ctx.World.Entities(), e)
			worldCollider := ecs.GetComponent[*components.WorldCollider](ctx.World.Entities(), e)

			if position.Y <= 0 {
				// top collision
				if worldCollider.Bounce {
					velocity.Y *= -1
				} else {
					position.Y = 0
				}

				ctx.EventWriter(events.WorldBoundaryEvent{}).SendOnce(
					events.WorldBoundaryEvent{
						Side: events.CollisionTop,
					},
				)
			} else if position.Y+asset.Handle.Height() >= worldBoundaries.Height {
				// bottom collision
				if worldCollider.Bounce {
					velocity.Y *= -1
				} else {
					position.Y = worldBoundaries.Height - asset.Handle.Height()
				}
				ctx.EventWriter(events.WorldBoundaryEvent{}).SendOnce(
					events.WorldBoundaryEvent{
						Side: events.CollisionBottom,
					},
				)
			}

			if position.X <= 0 {
				// left collision
				if worldCollider.Bounce {
					velocity.X *= -1
				} else {
					position.X = 0
				}
				ctx.EventWriter(events.WorldBoundaryEvent{}).SendOnce(
					events.WorldBoundaryEvent{
						Side: events.CollisionLeft,
					},
				)
			} else if position.X+asset.Handle.Width() >= worldBoundaries.Width {
				// right collision
				if worldCollider.Bounce {
					velocity.X *= -1
				} else {
					position.X = worldBoundaries.Width - asset.Handle.Width()
				}
				ctx.EventWriter(events.WorldBoundaryEvent{}).SendOnce(
					events.WorldBoundaryEvent{
						Side: events.CollisionRight,
					},
				)
			}
		}
	}
}
