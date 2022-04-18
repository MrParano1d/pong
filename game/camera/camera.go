package camera

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl"
	"github.com/mrparano1d/pong/opengl/camera"
)

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			config := ecs.GetResource[*opengl.PluginConfig](resourceMap)
			cam := camera.NewCamera(90.0, float32(config.Width), float32(config.Height))
			cam.Update()
			ecs.AddResource[*camera.Camera](resourceMap, cam)
		})
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {
		cam := ecs.GetResource[*camera.Camera](ctx.Resources)
		cam.Update()
	}
}
