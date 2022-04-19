package game

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/game/assets"
	"github.com/mrparano1d/pong/opengl"
)

type Plugin struct {
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStartUpSystemToStage(opengl.StagePrepare, assets.Setup())
	app.AddSystemToStage(opengl.StageRender, assets.System())
}
