package opengl

import (
	"github.com/mrparano1d/ecs"
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
		window := ecs.GetResource[*WindowRes](ctx.Resources).Handle
		window.SwapBuffers()
	})

	return s
}

func (p *CleanupStage) Name() string {
	return StageCleanup
}
