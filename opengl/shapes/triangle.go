package shapes

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/pong/opengl/types"
	"github.com/mrparano1d/pong/opengl/utils"
	"image/color"
)

type Triangle struct {
	vertices []types.Vertex

	shape *Shape
}

func NewTriangle(left, right, top mgl32.Vec3, color color.Color) Triangle {
	return Triangle{
		vertices: []types.Vertex{
			utils.ToVertex(left, color),
			utils.ToVertex(top, color),
			utils.ToVertex(right, color),
		},
	}
}

func (t *Triangle) Vertices() []types.Vertex {
	return t.vertices
}

func (t *Triangle) Indices() []uint32 {
	return []uint32{0, 1, 2}
}

func (t *Triangle) Create() error {
	shape, err := CreateShape(t.Vertices(), t.Indices())
	if err != nil {
		return fmt.Errorf("failed to triangle create shape: %v", err)
	}
	t.shape = shape

	return nil
}

func (t *Triangle) Bind() {
	t.shape.Bind()
}

func (t *Triangle) Unbind() {
	t.shape.Unbind()
}

func (t *Triangle) Draw() {
	gl.DrawElements(gl.TRIANGLES, int32(len(t.Indices())), gl.UNSIGNED_INT, nil)
}
