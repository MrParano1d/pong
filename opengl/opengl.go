package opengl

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"github.com/mrparano1d/pong/opengl/camera"
	"github.com/mrparano1d/pong/opengl/events"
	"github.com/mrparano1d/pong/opengl/time"
)

type PluginConfig struct {
	Title          string
	Width          int
	Height         int
	ShowWireframes bool
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

	// Events
	app.AddEvent(func(eventMap ecs.EventMap) {
		ecs.AddEvent[events.WindowResize](eventMap)
	})

	// Config
	app.AddStartUpSystemToStage(core.StageFirst, func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[*PluginConfig](resourceMap, p.config)
		})
	})

	// Time
	app.AddStartUpSystemToStage(core.StageFirst, time.Setup())
	app.AddSystemToStage(core.StageFirst, time.System())

	// Stages
	app.AddStageAfter(core.StagePostUpdate, NewRenderStage())
	app.AddStageBefore(StageRender, NewPrepareStage())
	app.AddStageBefore(StageRender, NewQueueStage())
	app.AddStageAfter(StageRender, NewCleanupStage())
	app.AddSystemToStage(core.StagePreUpdate, ClearSystem())

	// Camera
	app.AddStartUpSystemToStage(StagePrepare, camera.Setup())
	app.AddSystemToStage(StagePrepare, camera.System())

}
