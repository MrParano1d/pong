package core

import "github.com/mrparano1d/ecs"

type Plugin struct {
	environment string
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin(environment string) *Plugin {
	return &Plugin{
		environment: environment,
	}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStageBefore(ecs.StageUpdate, NewFirstStage(p.environment))
	app.AddStageAfter(ecs.StageUpdate, NewLastStage())
	app.AddStageBefore(ecs.StageUpdate, NewPreUpdateStage())
	app.AddStageAfter(ecs.StageUpdate, NewPostUpdateStage())
}
