package camera

import "github.com/go-gl/mathgl/mgl32"

type Projection interface {
	Matrix() mgl32.Mat4
	Update(width float32, height float32)
	DepthCalculation() DepthCalculation
	Far() float32
}
