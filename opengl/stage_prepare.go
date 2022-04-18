package opengl

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mrparano1d/ecs"
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
			window, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
			if err != nil {
				panic(err)
			}
			window.MakeContextCurrent()

			// Initialize Glow (go function bindings)
			if err := gl.Init(); err != nil {
				panic(err)
			}

			window.SetKeyCallback(keyCallback)

			gl.ClearColor(0.0, 0.0, 0.0, 1.0)

			ecs.AddResource[*WindowRes](resourceMap, &WindowRes{Handle: window})
			logger := ecs.GetResource[*zap.Logger](resourceMap)
			logger.Debug("OpenGL Version", zap.String("version", gl.GoStr(gl.GetString(gl.VERSION))))
		})
	})

	s.AddStartUpSystem(DebugSystem())

	s.AddStartUpSystem(func(commands ecs.Commands) {
		// TODO load assets
	})

	return s
}

func (p *PrepareStage) Name() string {
	return StagePrepare
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	// When a user presses the escape key, we set the WindowShouldClose property to true,
	// which closes the application
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
