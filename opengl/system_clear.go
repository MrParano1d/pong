package opengl

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/window"
)

func ClearSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		w := ecs.GetResource[*window.Resource](ctx.Resources).Handle

		if w.ShouldClose() {
			ctx.EventWriter(ecs.AppExitEvent{}).Send(ecs.AppExitEvent{})
		}

		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	}
}
