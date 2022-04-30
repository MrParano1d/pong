package systems

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/assets"
	"github.com/mrparano1d/pong/ebiten_plugin/events"
	gameevents "github.com/mrparano1d/pong/game/events"
	"github.com/mrparano1d/pong/game/resources"
	"strconv"
)

func GameSetup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(
			func(resourceMap ecs.ResourceMap) {
				ecs.AddResource[*resources.GameState](resourceMap, resources.NewGameState())
			},
		)
	}
}

func ScoreSystem() ecs.System {
	return func(ctx ecs.SystemContext) {
		reader := ctx.EventReader(events.WorldBoundaryEvent{})
		gameState := ecs.GetResource[*resources.GameState](ctx.Resources)
		assetManager := ecs.GetResource[*assets.Manager](ctx.Resources)

		for reader.Next() {
			ev := reader.Read().(events.WorldBoundaryEvent)

			switch ev.Side {
			case events.CollisionLeft:
				gameState.Player2Scored()
				ctx.EventWriter(gameevents.ScoredEvent{}).Send(gameevents.ScoredEvent{})
			case events.CollisionRight:
				gameState.Player1Scored()
				ctx.EventWriter(gameevents.ScoredEvent{}).Send(gameevents.ScoredEvent{})
			}
		}

		score1 := assetManager.Get("score1")
		score1.(*assets.Text).Text = strconv.Itoa(gameState.Player1Score)
		score2 := assetManager.Get("score2")
		score2.(*assets.Text).Text = strconv.Itoa(gameState.Player2Score)
	}
}
