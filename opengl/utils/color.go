package utils

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/pong/opengl/types"
	"image/color"
)

func ToOpenGLColor(c color.Color) (float32, float32, float32, float32) {
	r, g, b, a := c.RGBA()

	return float32(r) / 255.0, float32(g) / 255.0, float32(b) / 255.0, float32(a) / 255.0
}

func ToVertex(point mgl32.Vec3, color color.Color) types.Vertex {
	r, g, b, a := ToOpenGLColor(color)
	return types.Vertex{
		X: point.X(),
		Y: point.Y(),
		Z: point.Z(),

		R: r,
		G: g,
		B: b,
		A: a,
	}
}
