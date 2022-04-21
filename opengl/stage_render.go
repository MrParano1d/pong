package opengl

import (
	"github.com/mrparano1d/ecs"
	"runtime"
)

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()
}

const (
	StageRender = "render"
)

type RenderStage struct {
	ecs.Stage
}

var _ ecs.Stage = &RenderStage{}

func (RenderStage) Name() string {
	return StageRender
}

func NewRenderStage() *RenderStage {
	return &RenderStage{
		Stage: ecs.NewDefaultStage(),
	}
}
