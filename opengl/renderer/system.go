package renderer

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/camera"
	"github.com/mrparano1d/pong/opengl/cmd"
	"github.com/mrparano1d/pong/opengl/components"
	"sort"
)

func Setup2D() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[*Renderer](resourceMap, NewRenderer())
			ecs.AddResource[*RenderQueueResource](resourceMap, &RenderQueueResource{Queue: RenderQueue{}})
		})
		commands.Add(&cmd.AssetCommand{})
	}
}

type QueueItem2D struct {
	entity ecs.Entity
	order  float32
}

func Queue2D() ecs.System {
	return func(ctx ecs.SystemContext) {
		queueRes := ecs.GetResource[*RenderQueueResource](ctx.Resources)
		q := ecs.NewQuery(ctx.World)

		var order []QueueItem2D
		for _, entity := range q.Find(ecs.NewFilter(ecs.WithComponentFilter(&components.Asset{}, &components.Position{}))) {
			position := ecs.GetComponent[*components.Position](ctx.World.Entities(), entity)
			order = append(order, QueueItem2D{entity: entity, order: position.Y})
		}

		sort.SliceStable(order, func(i, j int) bool {
			return order[i].order > order[j].order
		})

		queue := make([]ecs.Entity, len(order))

		for i, o := range order {
			queue[i] = o.entity
		}

		queueRes.Queue.Replace(queue...)
	}
}

func System2D() ecs.System {
	return func(ctx ecs.SystemContext) {
		queueRes := ecs.GetResource[*RenderQueueResource](ctx.Resources)
		renderer := ecs.GetResource[*Renderer](ctx.Resources)
		cam := ecs.GetResource[*camera.Cameras](ctx.Resources)
		for _, entity := range queueRes.Queue {
			asset := ecs.GetComponent[*components.Asset](ctx.World.Entities(), entity)
			position := ecs.GetComponent[*components.Position](ctx.World.Entities(), entity)
			transform := ecs.GetComponent[*components.Transform2D](ctx.World.Entities(), entity)

			model := mgl32.Translate3D(0, 0, 0)

			// position
			model = model.Mul4(mgl32.Translate3D(position.X, position.Y, 0.0))
			if transform != nil {
				// move origin of rotation to center of quad
				model = model.Mul4(mgl32.Translate3D(0.5*transform.Scale.X(), 0.5*transform.Scale.Y(), 0.0))
				// rotation
				model = model.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(transform.Rotation)))
				// move origin back
				model = model.Mul4(mgl32.Translate3D(-0.5*transform.Scale.X(), -0.5*transform.Scale.Y(), 0.0))
				// scale
				model = model.Mul4(mgl32.Scale3D(transform.Scale.X(), transform.Scale.Y(), 1.0))
			}

			renderer.DrawAsset(cam.Orthographic.ViewProjection().Mul4(model), asset.Handle)
		}
	}
}
