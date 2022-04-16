package fps

import (
	"fmt"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl"
	"time"
)

type Plugin struct {
	lastUpdate int64
	count      int
	currentFPS float64
}

var _ ecs.Plugin = &Plugin{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Build(app *ecs.App) {
	app.AddStartUpSystemToStage(opengl.StageRender, func(commands *ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[*TickerRes](resourceMap, &TickerRes{Ticker: time.NewTicker(1 * time.Second)})
		})
	})

	app.AddSystemToStage(opengl.StageRender, func(ctx *ecs.SystemContext) {
		p.count++
		now := int64(time.Since(ctx.Time().Startup()))
		defer func() {
			p.lastUpdate = now
		}()

		if now < p.lastUpdate {
			panic("lastUpdate must be older than now")
		}

		ticker := ecs.GetResource[*TickerRes](ctx.Resources)

		select {
		case <-ticker.Ticker.C:
			fmt.Printf("Current FPS: %.2f\n", float64(p.count)*float64(time.Second)/float64(now-p.lastUpdate))
		default:
		}

		p.count = 0
	})

}
