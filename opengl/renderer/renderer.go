package renderer

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/pong/opengl/assets"
	"github.com/mrparano1d/pong/opengl/shader"
)

type Shaders struct {
	Basic *shader.Shader
}

type Renderer struct {
	shaders Shaders
}

func NewRenderer() *Renderer {
	s := shader.NewShader("./opengl/shaders/basic.vs", "./opengl/shaders/basic.fs", false)
	if err := s.Create(); err != nil {
		panic(fmt.Errorf("failed to create basic shader: %v", err))
	}
	s.Bind()

	return &Renderer{
		shaders: Shaders{
			Basic: s,
		},
	}
}

func (r *Renderer) DrawAsset(mvp mgl32.Mat4, asset assets.Asset) {
	asset.Bind()
	gl.UniformMatrix4fv(r.shaders.Basic.ModelViewProjMatrixLocation, 1, false, &mvp[0])
	asset.Draw()
	asset.Unbind()
}
