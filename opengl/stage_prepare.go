package opengl

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/cmd"
	"github.com/mrparano1d/pong/opengl/events"
	"github.com/mrparano1d/pong/opengl/window"
	"go.uber.org/zap"
	"log"
)

const (
	StagePrepare = "prepare"
)

type PrepareStage struct {
	ecs.Stage
}

var _ ecs.Stage = &PrepareStage{}

func NewPrepareStage() *PrepareStage {
	s := &PrepareStage{
		Stage: ecs.NewDefaultStage(),
	}

	s.AddStartUpSystem(func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			config := ecs.GetResource[*PluginConfig](resourceMap)

			if err := glfw.Init(); err != nil {
				log.Fatalln("failed to initialize glfw:", err)
			}
			//defer glfw.Terminate()

			glfw.WindowHint(glfw.Resizable, glfw.False)
			glfw.WindowHint(glfw.ContextVersionMajor, 4)
			glfw.WindowHint(glfw.ContextVersionMinor, 1)
			glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
			glfwWindow, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
			if err != nil {
				panic(err)
			}
			glfwWindow.MakeContextCurrent()

			// Initialize Glow (go function bindings)
			if err := gl.Init(); err != nil {
				panic(err)
			}

			glfwWindow.SetKeyCallback(keyCallback(config))

			gl.ClearColor(0.0, 0.0, 0.0, 1.0)

			ecs.AddResource[*window.Resource](resourceMap, &window.Resource{Handle: glfwWindow, Width: float32(config.Width), Height: float32(config.Height)})
			logger := ecs.GetResource[*zap.Logger](resourceMap)
			logger.Debug("OpenGL Version", zap.String("version", gl.GoStr(gl.GetString(gl.VERSION))))
		})

		commands.Add(&cmd.WindowCommand{
			ResizeCallback: func(world *ecs.World) {
				windowRes := ecs.GetResource[*window.Resource](world.Resources())
				windowRes.Handle.SetSizeCallback(func(w *glfw.Window, width int, height int) {
					world.Events()[events.WindowResize{}].Send(events.WindowResize{Width: width, Height: height})
					windowRes.Width = float32(width)
					windowRes.Height = float32(height)
				})
			},
		})
	})

	s.AddStartUpSystem(DebugSystem())

	s.AddSystem(func(ctx ecs.SystemContext) {
		config := ecs.GetResource[*PluginConfig](ctx.Resources)
		if config.ShowWireframes {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		} else {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		}
	})

	return s
}

func (p *PrepareStage) Name() string {
	return StagePrepare
}

func keyCallback(config *PluginConfig) func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	return func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

		if key == glfw.KeyW && action == glfw.Press {
			config.ShowWireframes = !config.ShowWireframes
		}

		// When a user presses the escape key, we set the WindowShouldClose property to true,
		// which closes the application
		if key == glfw.KeyEscape && action == glfw.Press {
			window.SetShouldClose(true)
		}
	}
}
