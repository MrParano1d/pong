package ebiten_plugin

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/internal/systems"
)

const (
	StagePrepare = "prepare"
)

type PrepareStage struct {
	ecs.Stage
}

var _ ecs.Stage = &PrepareStage{}

func NewPrepareStage() *PrepareStage {
	s := &PrepareStage{
		Stage: ecs.NewDefaultStage(),
	}

	s.AddStartUpSystem(systems.PrepareSetup())

	return s
}

func (p *PrepareStage) Name() string {
	return StagePrepare
}
