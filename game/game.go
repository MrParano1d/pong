package game

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"github.com/mrparano1d/pong/game/assets"
	"github.com/mrparano1d/pong/opengl"
	"github.com/mrparano1d/pong/opengl/renderer"
)

type Plugin struct {
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStartUpSystemToStage(opengl.StagePrepare, assets.Setup())
	app.AddStartUpSystemToStage(opengl.StagePrepare, renderer.Setup2D())
	app.AddSystemToStage(opengl.StageQueue, renderer.Queue2D())
	app.AddSystemToStage(core.StageUpdate, assets.System())
	app.AddSystemToStage(opengl.StageRender, renderer.System2D())
}
