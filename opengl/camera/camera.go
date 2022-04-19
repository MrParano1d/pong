package camera

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position   mgl32.Vec3
	projection Projection
	view       mgl32.Mat4
	viewProj   mgl32.Mat4
}

func NewCamera(projection Projection) *Camera {

	c := &Camera{
		projection: projection,
		view:       mgl32.Translate3D(0, 0, 0),
		position:   mgl32.Vec3{0.0, 0.0, 0.0},
	}

	c.Update()

	return c
}

func (c *Camera) Projection() Projection {
	return c.projection
}

func (c *Camera) Update() {
	c.viewProj = c.projection.Matrix().Mul4(c.view)
}

func (c *Camera) Translate(v mgl32.Vec3) {
	c.position = c.position.Add(v)
	c.view = c.view.Mul4(mgl32.Translate3D(v.X()*-1.0, v.Y()*-1.0, v.Z()*-1))
}

func (c *Camera) ViewProjection() mgl32.Mat4 {
	return c.viewProj
}
