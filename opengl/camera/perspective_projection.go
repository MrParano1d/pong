package camera

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type PerspectiveProjection struct {
	fov         float32
	aspectRatio float32
	near        float32
	far         float32
}

var _ Projection = &PerspectiveProjection{}

type PerspectiveProjectionOption func(p *PerspectiveProjection)

func WithPerspectiveProjection(pp PerspectiveProjection) PerspectiveProjectionOption {
	return func(p *PerspectiveProjection) {
		*p = pp
	}
}

func NewPerspectiveProjection(opts ...PerspectiveProjectionOption) *PerspectiveProjection {
	p := &PerspectiveProjection{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func NewDefaultPerspectiveProjection() *PerspectiveProjection {
	return &PerspectiveProjection{
		fov:         math.Pi / 4,
		near:        0.1,
		far:         1000.0,
		aspectRatio: 1.0,
	}
}

func (p *PerspectiveProjection) Matrix() mgl32.Mat4 {
	return mgl32.Perspective(p.fov, p.aspectRatio, p.near, p.far)
}

func (p *PerspectiveProjection) Update(width float32, height float32) {
	p.aspectRatio = width / height
}

func (p *PerspectiveProjection) DepthCalculation() DepthCalculation {
	return DepthCalculationDistance
}

func (p *PerspectiveProjection) Far() float32 {
	return p.far
}
