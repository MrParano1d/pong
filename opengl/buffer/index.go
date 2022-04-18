package buffer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"reflect"
)

type IndexBuffer struct {
	bufferId uint32
}

func NewIndexBuffer(data []uint32) *IndexBuffer {
	b := &IndexBuffer{}

	elementSize := int(reflect.ValueOf(data).Index(0).Type().Size())

	gl.GenBuffers(1, &b.bufferId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, b.bufferId)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data)*elementSize, gl.Ptr(data), gl.STATIC_DRAW)

	return b
}

func (b *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, b.bufferId)
}

func (b *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (b *IndexBuffer) Delete() {
	gl.DeleteBuffers(1, &b.bufferId)
}
