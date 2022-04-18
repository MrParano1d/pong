package time

import "github.com/go-gl/glfw/v3.3/glfw"

type Resource struct {
	delta       float64
	lastCounter float64
}

func (t *Resource) Delta() float64 {
	return t.delta
}

func (t *Resource) Update() {
	counter := glfw.GetTime()
	t.delta = counter - t.lastCounter
	t.lastCounter = counter
}
