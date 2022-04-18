package shapes

import (
	"fmt"
	"github.com/mrparano1d/pong/opengl/buffer"
	"github.com/mrparano1d/pong/opengl/shader"
	"github.com/mrparano1d/pong/opengl/types"
)

type Shape struct {
	indexBuffer  *buffer.IndexBuffer
	vertexBuffer *buffer.VertexBuffer
	shader       *shader.Shader

	modelViewProjMatrixLocation int32
}

func CreateShape(vertices []types.Vertex, indices []uint32) (*Shape, error) {

	s := &Shape{}

	s.indexBuffer = buffer.NewIndexBuffer(indices)
	s.vertexBuffer = buffer.NewVertexBuffer(vertices)
	s.vertexBuffer.Unbind()

	// @TODO set cleanup to true in production
	s.shader = shader.NewShader("./opengl/shaders/basic.vs", "./opengl/shaders/basic.fs", false)
	if err := s.shader.Create(); err != nil {
		return nil, fmt.Errorf("failed to create triangle shader: %v", err)
	}
	s.shader.Bind()

	s.modelViewProjMatrixLocation = s.shader.UniformLocation("u_modelViewProj")

	return s, nil
}

func (s *Shape) Bind() {
	s.vertexBuffer.Bind()
	s.indexBuffer.Bind()
}

func (s *Shape) Unbind() {
	s.indexBuffer.Unbind()
	s.vertexBuffer.Unbind()
}
