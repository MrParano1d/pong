package camera

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position   mgl32.Vec3
	projection mgl32.Mat4
	view       mgl32.Mat4
	viewProj   mgl32.Mat4
}

func NewCamera(fov float32, width float32, height float32) *Camera {
	c := &Camera{
		projection: mgl32.Perspective(fov/2.0, width/height, 0.1, 1000.0),
		view:       mgl32.Translate3D(0, 0, -2.0),
		position:   mgl32.Vec3{0.0, 0.0, 0.0},
	}

	c.Update()

	return c
}

func (c *Camera) Update() {
	c.viewProj = c.projection.Mul4(c.view)
}

func (c *Camera) Translate(v mgl32.Vec3) {
	c.position = c.position.Add(v)
	translation := mgl32.Translate3D(v.X()*-1.0, v.Y()*-1.0, v.Z()*-1)
	c.view = c.view.Mul4(translation)
}

func (c *Camera) ViewProjection() mgl32.Mat4 {
	return c.viewProj
}
