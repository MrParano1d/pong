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
		projection: mgl32.Ortho(0.0, width, height, 0.0, -1, 1),
		//view:       mgl32.LookAt(cameraPos.X(), cameraPos.Y(), cameraPos.Z(), cameraCenter.X(), cameraCenter.Y(), cameraCenter.Z(), cameraUp.X(), cameraUp.Y(), cameraUp.Z()),
		position: mgl32.Vec3{0.0, 0.0, 0.0},
	}

	c.Update()

	return c
}

func (c *Camera) Update() {
	c.viewProj = c.projection
}

func (c *Camera) Translate(v mgl32.Vec3) {
	c.position = c.position.Add(v)
	c.view = c.view.Mul4(mgl32.Translate3D(v.X()*-1.0, v.Y()*-1.0, v.Z()*-1))
}

func (c *Camera) ViewProjection() mgl32.Mat4 {
	return c.viewProj
}
