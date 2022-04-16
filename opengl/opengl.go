package opengl

import "github.com/mrparano1d/ecs"

type Plugin struct {
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStageAfter(ecs.StageUpdate, StageRender, NewRenderStage())
}
