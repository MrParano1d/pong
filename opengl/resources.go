package opengl

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mrparano1d/pong/gfx"
)

type WindowRes struct {
	Handle *glfw.Window
}

type VAORes uint32

type Texture0Res struct {
	Texture *gfx.Texture
}
type Texture1Res struct {
	Texture *gfx.Texture
}

type ShaderProgramRes struct {
	Handle *gfx.Program
}
