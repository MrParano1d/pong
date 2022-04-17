package opengl

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/core"
)

type Plugin struct {
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStageAfter(core.StagePostUpdate, NewRenderStage())
	app.AddStageBefore(StageRender, NewPrepareStage())
	app.AddStageAfter(StageRender, NewCleanupStage())
	app.AddSystemToStage(core.StagePreUpdate, ClearSystem())
}
