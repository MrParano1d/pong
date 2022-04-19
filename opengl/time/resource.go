package time

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"time"
)

type Resource struct {
	delta       float64
	lastCounter float64
	startup     time.Time
}

func (t *Resource) Startup() time.Time {
	return t.startup
}

func (t *Resource) Delta() float64 {
	return t.delta
}

func (t *Resource) Update() {
	counter := glfw.GetTime()
	t.delta = counter - t.lastCounter
	t.lastCounter = counter
}
