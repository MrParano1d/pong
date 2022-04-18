package shader

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/mrparano1d/pong/opengl/utils"
	"io/ioutil"
)

type Shader struct {
	shaderId               uint32
	vertexShaderFilepath   string
	fragmentShaderFilepath string
	cleanup                bool
}

func NewShader(vertexShaderFilepath string, fragmentShaderFilepath string, cleanup bool) *Shader {
	return &Shader{
		vertexShaderFilepath:   vertexShaderFilepath,
		fragmentShaderFilepath: fragmentShaderFilepath,
		cleanup:                cleanup,
	}
}

func (s *Shader) Create() error {
	s.shaderId = gl.CreateProgram()

	vsBytes, err := ioutil.ReadFile(s.vertexShaderFilepath)
	if err != nil {
		return fmt.Errorf("shader failed to load vertex shader %s: %v", s.vertexShaderFilepath, err)
	}
	fsBytes, err := ioutil.ReadFile(s.fragmentShaderFilepath)
	if err != nil {
		return fmt.Errorf("shader failed to load fragment shader %s: %v", s.fragmentShaderFilepath, err)
	}

	vs, err := s.compile(vsBytes, gl.VERTEX_SHADER)
	if err != nil {
		return fmt.Errorf("failed to compile vertex shader: %v", err)
	}
	fs, err := s.compile(fsBytes, gl.FRAGMENT_SHADER)
	if err != nil {
		return fmt.Errorf("failed to compile fragment shader: %v", err)
	}

	return s.link([2]uint32{vs, fs})
}

func (s *Shader) link(shaders [2]uint32) error {
	for _, sh := range shaders {
		gl.AttachShader(s.shaderId, sh)
	}
	gl.LinkProgram(s.shaderId)

	if err := utils.CheckGlError(s.shaderId, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog); err != nil {
		return fmt.Errorf("failed to link shader program: %v", err)
	}

	if s.cleanup {
		for _, sh := range shaders {
			gl.DetachShader(s.shaderId, sh)
			gl.DeleteShader(sh)
		}
	}

	return nil
}

func (s *Shader) compile(source []byte, shaderType uint32) (uint32, error) {
	id := gl.CreateShader(shaderType)

	shaderChars, freeFn := gl.Strs(string(source) + "\x00")
	defer freeFn()

	gl.ShaderSource(id, 1, shaderChars, nil)
	gl.CompileShader(id)

	if err := utils.CheckGlError(id, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog); err != nil {
		return 0, fmt.Errorf("failed to compile shader source: %v", err)
	}

	return id, nil
}

func (s *Shader) UniformLocation(uniformName string) int32 {
	return gl.GetUniformLocation(s.shaderId, gl.Str(uniformName+"\x00"))
}

func (s *Shader) ShaderID() uint32 {
	return s.shaderId
}

func (s *Shader) Bind() {
	gl.UseProgram(s.shaderId)
}

func (s *Shader) Unbind() {
	gl.UseProgram(0)
}
