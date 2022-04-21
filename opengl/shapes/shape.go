package shapes

import (
	"github.com/mrparano1d/pong/opengl/assets"
	"github.com/mrparano1d/pong/opengl/buffer"
	"github.com/mrparano1d/pong/opengl/types"
)

type Shape struct {
	indexBuffer  *buffer.IndexBuffer
	vertexBuffer *buffer.VertexBuffer
}

var _ assets.Asset = &Shape{}

func CreateShape(vertices []types.Vertex, indices []uint32) (*Shape, error) {

	s := &Shape{}

	s.indexBuffer = buffer.NewIndexBuffer(indices)
	s.vertexBuffer = buffer.NewVertexBuffer(vertices)
	s.vertexBuffer.Unbind()

	return s, nil
}

func (s *Shape) Create() error {
	return nil
}

func (s *Shape) Width() float32 {
	return 0.0
}

func (s *Shape) Height() float32 {
	return 0.0
}

func (s *Shape) Bind() {
	s.vertexBuffer.Bind()
	s.indexBuffer.Bind()
}

func (s *Shape) Unbind() {
	s.indexBuffer.Unbind()
	s.vertexBuffer.Unbind()
}

func (s *Shape) Draw() {

}
