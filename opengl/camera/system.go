package camera

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/events"
	"github.com/mrparano1d/pong/opengl/window"
)

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			w := ecs.GetResource[*window.Resource](resourceMap)

			oProj := NewDefaultOrthographicProjection(WithOrthographicWindowOrigin(WindowOriginBottomLeft))
			oProj.Update(w.Width, w.Height)

			pProj := NewDefaultPerspectiveProjection()
			pProj.Update(w.Width, w.Height)

			ecs.AddResource[*Projections](resourceMap, &Projections{
				Orthographic: oProj,
				Perspective:  pProj,
			})

			ecs.AddResource[*Cameras](resourceMap, &Cameras{
				Orthographic: NewCamera(oProj),
				Perspective:  NewCamera(pProj),
			})
		})
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {

		reader := ctx.EventReader(events.WindowResize{})
		for reader.Next() {
			event := reader.Read().(events.WindowResize)

			cameras := ecs.GetResource[*Cameras](ctx.Resources)

			projections := ecs.GetResource[*Projections](ctx.Resources)

			projections.Orthographic.Update(float32(event.Width), float32(event.Height))
			projections.Perspective.Update(float32(event.Width), float32(event.Height))

			cameras.Orthographic.Update()
			cameras.Perspective.Update()
		}
	}
}
