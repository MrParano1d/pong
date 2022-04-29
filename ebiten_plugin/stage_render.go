package ebiten_plugin

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/internal/systems"
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
	s := &RenderStage{
		Stage: ecs.NewDefaultStage(ecs.WithStageLabel(ecs.LabelRender)),
	}

	s.AddSystem(systems.RenderSystem())

	return s
}
