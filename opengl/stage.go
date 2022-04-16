package opengl

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/gfx"
	"log"
	"runtime"
	"unsafe"
)

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()
}

const (
	StageRender  = "render"
	windowWidth  = 800
	windowHeight = 600
)

type RenderStage struct {
	ecs.Stage
}

var _ ecs.Stage = &RenderStage{}

func NewRenderStage() *RenderStage {
	s := &RenderStage{
		Stage: ecs.NewDefaultStage(),
	}

	s.AddStartUpSystem(func(commands *ecs.Commands) {
		if err := glfw.Init(); err != nil {
			log.Fatalln("failed to inifitialize glfw:", err)
		}
		//defer glfw.Terminate()

		glfw.WindowHint(glfw.Resizable, glfw.False)
		glfw.WindowHint(glfw.ContextVersionMajor, 4)
		glfw.WindowHint(glfw.ContextVersionMinor, 1)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		window, err := glfw.CreateWindow(windowWidth, windowHeight, "basic textures", nil, nil)
		if err != nil {
			panic(err)
		}
		window.MakeContextCurrent()

		// Initialize Glow (go function bindings)
		if err := gl.Init(); err != nil {
			panic(err)
		}

		window.SetKeyCallback(keyCallback)

		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[*WindowRes](resourceMap, &WindowRes{Handle: window})
		})
	})
	s.AddStartUpSystem(func(commands *ecs.Commands) {
		// the linked shader program determines how the data will be rendered
		vertShader, err := gfx.NewShaderFromFile("./shaders/basic.vert", gl.VERTEX_SHADER)
		if err != nil {
			log.Fatal(err)
		}

		fragShader, err := gfx.NewShaderFromFile("./shaders/basic.frag", gl.FRAGMENT_SHADER)
		if err != nil {
			log.Fatal(err)
		}

		shaderProgram, err := gfx.NewProgram(vertShader, fragShader)
		if err != nil {
			log.Fatal(err)
		}
		//defer shaderProgram.Delete()

		vertices := []float32{
			// top left
			-0.75, 0.75, 0.0, // position
			1.0, 0.0, 0.0, // Color
			1.0, 0.0, // texture coordinates

			// top right
			0.75, 0.75, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0,

			// bottom right
			0.75, -0.75, 0.0,
			0.0, 0.0, 1.0,
			0.0, 1.0,

			// bottom left
			-0.75, -0.75, 0.0,
			1.0, 1.0, 1.0,
			1.0, 1.0,
		}

		indices := []uint32{
			// rectangle
			0, 1, 2, // top triangle
			0, 2, 3, // bottom triangle
		}

		VAO := createVAO(vertices, indices)

		texture0, err := gfx.NewTextureFromFile("./images/RTS_Crate.png", gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
		if err != nil {
			panic(err.Error())
		}

		texture1, err := gfx.NewTextureFromFile("./images/trollface.png", gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
		if err != nil {
			panic(err.Error())
		}

		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[VAORes](resourceMap, VAORes(VAO))
			ecs.AddResource[*Texture0Res](resourceMap, &Texture0Res{Texture: texture0})
			ecs.AddResource[*Texture1Res](resourceMap, &Texture1Res{Texture: texture1})
			ecs.AddResource[*ShaderProgramRes](resourceMap, &ShaderProgramRes{Handle: shaderProgram})
		})
	})
	s.AddSystem(func(ctx *ecs.SystemContext) {
		window := ecs.GetResource[*WindowRes](ctx.Resources).Handle

		if window.ShouldClose() {
			ctx.Commands.Cancel()
		}

		texture0 := ecs.GetResource[*Texture0Res](ctx.Resources).Texture
		texture1 := ecs.GetResource[*Texture1Res](ctx.Resources).Texture
		shaderProgram := ecs.GetResource[*ShaderProgramRes](ctx.Resources).Handle
		VAO := uint32(ecs.GetResource[VAORes](ctx.Resources))

		glfw.PollEvents()

		// background color
		gl.ClearColor(0.2, 0.5, 0.5, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// draw vertices
		shaderProgram.Use()

		// set texture0 to uniform0 in the fragment shader
		texture0.Bind(gl.TEXTURE0)
		texture0.SetUniform(shaderProgram.GetUniformLocation("ourTexture0"))

		// set texture1 to uniform1 in the fragment shader
		texture1.Bind(gl.TEXTURE1)
		texture1.SetUniform(shaderProgram.GetUniformLocation("ourTexture1"))

		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
		gl.BindVertexArray(0)

		texture0.UnBind()
		texture1.UnBind()

		// end of draw loop

		// swap in the rendered buffer
		window.SwapBuffers()
	})

	return s
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	// When a user presses the escape key, we set the WindowShouldClose property to true,
	// which closes the application
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}

/*
 * Creates the Vertex Array Object for a triangle.
 */
func createVAO(vertices []float32, indices []uint32) uint32 {

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)

	var EBO uint32
	gl.GenBuffers(1, &EBO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// copy indices into element buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 3*4 + 3*4 + 2*4
	var offset int = 0

	// position
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, stride, uintptr(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// color
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, stride, uintptr(offset))
	gl.EnableVertexAttribArray(1)
	offset += 3 * 4

	// texture position
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, stride, uintptr(offset))
	gl.EnableVertexAttribArray(2)
	offset += 2 * 4

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}
