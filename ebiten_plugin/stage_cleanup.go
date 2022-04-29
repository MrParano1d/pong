package ebiten_plugin

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
	return s
}

func (p *CleanupStage) Name() string {
	return StageCleanup
}
