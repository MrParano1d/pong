package opengl

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
)

type PluginConfig struct {
	Title  string
	Width  int
	Height int
}

type Plugin struct {
	config *PluginConfig
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin(config *PluginConfig) *Plugin {
	return &Plugin{
		config: config,
	}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStartUpSystemToStage(core.StageFirst, func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[*PluginConfig](resourceMap, p.config)
		})
	})
	app.AddStageAfter(core.StagePostUpdate, NewRenderStage())
	app.AddStageBefore(StageRender, NewPrepareStage())
	app.AddStageAfter(StageRender, NewCleanupStage())
	app.AddSystemToStage(core.StagePreUpdate, ClearSystem())
}
