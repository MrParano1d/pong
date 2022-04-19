package camera

import (
	"github.com/go-gl/mathgl/mgl32"
)

type OrthographicProjection struct {
	left             float32
	right            float32
	bottom           float32
	top              float32
	near             float32
	far              float32
	windowOrigin     WindowOrigin
	scalingMode      ScalingMode
	scale            float32
	depthCalculation DepthCalculation
}

var _ Projection = &OrthographicProjection{}

type OrthographicProjectionOption func(p *OrthographicProjection)

func WithOrthographicWindowOrigin(windowOrigin WindowOrigin) OrthographicProjectionOption {
	return func(p *OrthographicProjection) {
		p.windowOrigin = windowOrigin
	}
}

func NewOrthographicProjection(
	left float32,
	right float32,
	bottom float32,
	top float32,
	near float32,
	far float32,
	windowOrigin WindowOrigin,
	scalingMode ScalingMode,
	scale float32,
	depthCalculation DepthCalculation,
) *OrthographicProjection {
	p := &OrthographicProjection{
		left:             left,
		right:            right,
		bottom:           bottom,
		top:              top,
		near:             near,
		far:              far,
		windowOrigin:     windowOrigin,
		scalingMode:      scalingMode,
		scale:            scale,
		depthCalculation: depthCalculation,
	}

	return p
}

func NewDefaultOrthographicProjection(opts ...OrthographicProjectionOption) *OrthographicProjection {
	p := &OrthographicProjection{
		left:             -1.0,
		right:            1.0,
		bottom:           -1.0,
		top:              1.0,
		near:             0.0,
		far:              1000.0,
		windowOrigin:     WindowOriginCenter,
		scalingMode:      ScalingModeWindowSize,
		scale:            1.0,
		depthCalculation: DepthCalculationDistance,
	}

	for _, opt := range opts {
		opt(p)
	}

	p.Update(60, 60)

	return p
}

func (p *OrthographicProjection) Matrix() mgl32.Mat4 {
	return mgl32.Ortho(p.left*p.scale, p.right*p.scale, p.bottom*p.scale, p.top*p.scale, p.far, p.near)
}

func (p *OrthographicProjection) Update(width float32, height float32) {
	if p.scalingMode == ScalingModeWindowSize && p.windowOrigin == WindowOriginCenter {
		halfWidth := width / 2.0
		halfHeight := height / 2.0

		p.left = -halfWidth
		p.right = halfWidth
		p.top = halfHeight
		p.bottom = -halfHeight
	} else if p.scalingMode == ScalingModeWindowSize && p.windowOrigin == WindowOriginBottomLeft {
		p.left = 0.0
		p.right = width
		p.top = height
		p.bottom = 0.0
	} else if p.scalingMode == ScalingModeFixedVertical && p.windowOrigin == WindowOriginCenter {
		aspectRatio := width / height
		p.left = -aspectRatio
		p.right = aspectRatio
		p.top = 1.0
		p.bottom = -1.0
	} else if p.scalingMode == ScalingModeFixedVertical && p.windowOrigin == WindowOriginBottomLeft {
		aspectRatio := width / height
		p.left = 0.0
		p.right = aspectRatio
		p.top = 1.0
		p.bottom = 0.0
	} else if p.scalingMode == ScalingModeFixedHorizontal && p.windowOrigin == WindowOriginCenter {
		aspectRatio := height / width
		p.left = -1.0
		p.right = 1.0
		p.top = aspectRatio
		p.bottom = -aspectRatio
	} else if p.scalingMode == ScalingModeFixedHorizontal && p.windowOrigin == WindowOriginBottomLeft {
		aspectRatio := height / width
		p.left = 0.0
		p.right = 1.0
		p.top = aspectRatio
		p.bottom = 0.0
	}
}

func (p *OrthographicProjection) DepthCalculation() DepthCalculation {
	return p.depthCalculation
}

func (p *OrthographicProjection) Far() float32 {
	return p.far
}
