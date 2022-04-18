package buffer

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/mrparano1d/pong/opengl/types"
	"reflect"
)

type VertexBuffer struct {
	vao      uint32
	bufferId uint32
}

func NewVertexBuffer(data []types.Vertex) *VertexBuffer {

	vertexSize := int(reflect.TypeOf(types.Vertex{}).Size())

	b := &VertexBuffer{}

	gl.GenVertexArrays(1, &b.vao)
	gl.BindVertexArray(b.vao)

	gl.GenBuffers(1, &b.bufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, b.bufferId)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*vertexSize, gl.Ptr(data), gl.STATIC_DRAW)

	xField, ok := reflect.TypeOf(types.Vertex{}).FieldByName("X")
	if !ok {
		panic(fmt.Errorf("%T doesn't contain a field named 'X'", types.Vertex{}))
	}
	rField, ok := reflect.TypeOf(types.Vertex{}).FieldByName("R")
	if !ok {
		panic(fmt.Errorf("%T doesn't contain a field named 'R'", types.Vertex{}))
	}

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, int32(vertexSize), xField.Offset)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 4, gl.FLOAT, false, int32(vertexSize), rField.Offset)

	gl.BindVertexArray(0)

	return b
}

func (b *VertexBuffer) Delete() {
	gl.DeleteBuffers(1, &b.bufferId)
}

func (b *VertexBuffer) Bind() {
	gl.BindVertexArray(b.vao)
}

func (b *VertexBuffer) Unbind() {
	gl.BindVertexArray(0)
}
