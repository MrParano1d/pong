package ebiten_plugin

import (
	"fmt"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"
	render "github.com/mrparano1d/pong/ebiten_plugin/internal/render/ebiten"
	"image/color"
)

type Plugin struct {
	renderer *render.Renderer
	app      *ecs.App
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin() *Plugin {
	r, err := render.CreateRenderer()
	if err != nil {
		panic(err)
	}
	return &Plugin{
		renderer: r,
	}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStageAfter(core.StagePostUpdate, NewRenderStage())
	app.AddStageBefore(StageRender, NewPrepareStage())
	app.AddStageAfter(StageRender, NewCleanupStage())

	p.app = app
}

func (p *Plugin) Run() error {

	p.app.SetupSystems()

	return p.renderer.Run(
		func(surface interfaces.Surface) error {
			ecs.AddResource[interfaces.Surface](p.app.World().Resources(), surface)

			surface.DrawText(
				fmt.Sprintf(
					"TPS: %0.2f / FPS: %0.2f",
					p.renderer.CurrentTPS(),
					p.renderer.CurrentFPS(),
				),
				color.RGBA{255, 255, 255, 255},
			)

			p.app.RunSystems(ecs.WithStageLabelFilter(ecs.LabelRender))
			return nil
		}, func() error {
			p.app.RunSystems(ecs.WithStageLabelFilter(ecs.LabelNone, ecs.LabelUpdate))
			return nil
		}, 512*2, 256*2, "Pong",
	)
}
