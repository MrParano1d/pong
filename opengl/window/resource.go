package window

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Resource struct {
	Handle *glfw.Window
	Width  float32
	Height float32
}

func (r *Resource) SetWindowSize(width, height float32) {
	r.Handle.SetSize(int(width), int(height))
	r.Handle.SetSizeLimits(int(width), int(height), int(width), int(height))
	gl.Viewport(0, 0, int32(width), int32(height))
}
