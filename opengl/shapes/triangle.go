package shapes

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/pong/opengl/buffer"
	"github.com/mrparano1d/pong/opengl/shader"
	"github.com/mrparano1d/pong/opengl/types"
	"github.com/mrparano1d/pong/opengl/utils"
	"image/color"
)

type Triangle struct {
	vertices []types.Vertex

	indexBuffer  *buffer.IndexBuffer
	vertexBuffer *buffer.VertexBuffer
	shader       *shader.Shader

	modelViewProjMatrixLocation int32
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
	t.indexBuffer = buffer.NewIndexBuffer(t.Indices())
	t.vertexBuffer = buffer.NewVertexBuffer(t.Vertices())
	t.vertexBuffer.Unbind()

	// @TODO set cleanup to true in production
	t.shader = shader.NewShader("./opengl/shaders/basic.vs", "./opengl/shaders/basic.fs", false)
	if err := t.shader.Create(); err != nil {
		return fmt.Errorf("failed to create triangle shader: %v", err)
	}
	t.shader.Bind()

	t.modelViewProjMatrixLocation = t.shader.UniformLocation("u_modelViewProj")

	return nil
}

func (t *Triangle) Bind() {
	t.vertexBuffer.Bind()
	t.indexBuffer.Bind()
}

func (t *Triangle) Unbind() {
	t.indexBuffer.Unbind()
	t.vertexBuffer.Unbind()
}

func (t *Triangle) Draw(modelViewProj mgl32.Mat4) {
	t.Bind()
	defer t.Unbind()
	gl.UniformMatrix4fv(t.modelViewProjMatrixLocation, 1, false, &modelViewProj[0])
	gl.DrawElements(gl.TRIANGLES, int32(len(t.Indices())), gl.UNSIGNED_INT, nil)
}
