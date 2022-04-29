package game

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"github.com/mrparano1d/pong/ebiten_plugin"
	"github.com/mrparano1d/pong/game/assets"
	"github.com/mrparano1d/pong/game/camera"
	"github.com/mrparano1d/pong/game/systems"
)

type Plugin struct {
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStartUpSystemToStage(ebiten_plugin.StagePrepare, camera.Setup())
	app.AddSystemToStage(core.StagePostUpdate, camera.System())

	app.AddSystemToStage(core.StageUpdate, systems.PlayerSystem())

	app.AddStartUpSystemToStage(ebiten_plugin.StagePrepare, assets.Setup())
	app.AddSystemToStage(ebiten_plugin.StageRender, assets.System())
}
