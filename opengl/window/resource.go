package window

import "github.com/go-gl/glfw/v3.3/glfw"

type Resource struct {
	Handle *glfw.Window
	Width  float32
	Height float32
}
