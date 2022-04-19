package opengl

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/window"
)

const (
	StageCleanup = "cleanup"
)

type CleanupStage struct {
	ecs.Stage
}

var _ ecs.Stage = &CleanupStage{}

func NewCleanupStage() *CleanupStage {
	s := &CleanupStage{
		Stage: ecs.NewDefaultStage(),
	}

	s.AddSystem(func(ctx ecs.SystemContext) {
		w := ecs.GetResource[*window.Resource](ctx.Resources).Handle
		w.SwapBuffers()
	})

	return s
}

func (p *CleanupStage) Name() string {
	return StageCleanup
}
