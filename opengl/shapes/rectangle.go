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

	Width  int
	Height int
	X      int
	Y      int

	shape *Shape
}

func NewRectangle(x, y, width, height int, color color.Color) Rectangle {

	posX := float32(x)
	posY := float32(y)
	widthF := float32(width)
	heightF := float32(height)

	return Rectangle{
		vertices: []types.Vertex{
			utils.ToVertex([3]float32{posX, posY, 0.0}, color),
			utils.ToVertex([3]float32{posX, posY + heightF, 0.0}, color),
			utils.ToVertex([3]float32{posX + widthF, posY, 0.0}, color),
			utils.ToVertex([3]float32{posX + widthF, posY + heightF, 0.0}, color),
		},
		Width:  width,
		Height: height,
		X:      x,
		Y:      y,
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
