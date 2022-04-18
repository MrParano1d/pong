package shapes

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/pong/opengl/types"
	"github.com/mrparano1d/pong/opengl/utils"
	"image/color"
)

type Rectangle struct {
	vertices []types.Vertex

	shape *Shape
}

func NewRectangle(bottomLeft, topLeft, topRight, bottomRight mgl32.Vec3, color color.Color) Rectangle {
	return Rectangle{
		vertices: []types.Vertex{
			utils.ToVertex(bottomLeft, color),
			utils.ToVertex(topLeft, color),
			utils.ToVertex(topRight, color),
			utils.ToVertex(bottomRight, color),
		},
	}
}

func (r *Rectangle) Vertices() []types.Vertex {
	return r.vertices
}

func (r *Rectangle) Indices() []uint32 {
	return []uint32{
		0, 1, 2,
		1, 2, 3,
	}
}

func (r *Rectangle) Create() error {
	shape, err := CreateShape(r.Vertices(), r.Indices())
	if err != nil {
		return fmt.Errorf("failed to create rectangle shape: %v", err)
	}
	r.shape = shape

	return nil
}

func (r *Rectangle) Bind() {
	r.shape.Bind()
}

func (r *Rectangle) Unbind() {
	r.shape.Unbind()
}

func (r *Rectangle) Draw(modelViewProj mgl32.Mat4) {
	r.Bind()
	defer r.Unbind()
	gl.UniformMatrix4fv(r.shape.modelViewProjMatrixLocation, 1, false, &modelViewProj[0])
	gl.DrawElements(gl.TRIANGLES, int32(len(r.Indices())), gl.UNSIGNED_INT, nil)
}
