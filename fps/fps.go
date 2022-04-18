package fps

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"go.uber.org/zap"
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

	app.AddStartUpSystemToStage(core.StageFirst, func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			config := ecs.GetResource[*core.ConfigRes](resourceMap)
			if config.Environment == core.EnvRelease {
				return
			}
			ecs.AddResource[*TickerRes](resourceMap, &TickerRes{Ticker: time.NewTicker(1 * time.Second)})
		})
	})

	app.AddSystemToStage(core.StageFirst, func(ctx ecs.SystemContext) {
		config := ecs.GetResource[*core.ConfigRes](ctx.Resources)
		if config.Environment == core.EnvRelease {
			return
		}

		gameTime := ecs.GetResource[*core.Time](ctx.Resources)

		p.count++
		now := int64(time.Since(gameTime.Startup()))
		defer func() {
			p.lastUpdate = now
		}()

		if now < p.lastUpdate {
			panic("lastUpdate must be older than now")
		}

		ticker := ecs.GetResource[*TickerRes](ctx.Resources)

		select {
		case <-ticker.Ticker.C:
			logger := ecs.GetResource[*zap.Logger](ctx.Resources)
			logger.Debug("FPS", zap.Float64("fps", float64(p.count)*float64(time.Second)/float64(now-p.lastUpdate)))
		default:
		}

		p.count = 0
	})
}
