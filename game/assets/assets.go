package assets

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/assets"
	ebitencomponents "github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/ebiten_plugin/enums"
	"github.com/mrparano1d/pong/game/components"
	"image/color"
)

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {

		player1 := assets.NewRectangle(4, 28, 0, 0, color.RGBA{255, 255, 255, 255}, nil)
		player2 := assets.NewRectangle(4, 28, 0, 0, color.RGBA{255, 255, 255, 255}, nil)

		middleline := assets.NewRectangle(2, 256, 0, 0, color.RGBA{255, 255, 255, 255}, nil)

		commands.Add(
			assets.NewCommand(nil).Add(
				"player1",
				player1,
			).Add(
				"player2",
				player2,
			).Add(
				"middleline",
				middleline,
			),
		)

		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: middleline}).Insert(
			&ebitencomponents.Position{
				X: (512 / 2) + 1,
				Y: 0,
			},
		)

		// Player 1
		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: player1}).Insert(
			&ebitencomponents.Position{
				X: 14,
				Y: 100,
			},
		).Insert(&components.Player{UpKey: enums.KeyW, DownKey: enums.KeyS}).Insert(
			&ebitencomponents.Velocity{
				X: 2,
				Y: 2,
				Z: 0,
			},
		)

		// Player 2
		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: player2}).Insert(
			&ebitencomponents.Position{
				X: 512 - 14,
				Y: 100,
			},
		).Insert(&components.Player{UpKey: enums.KeyUp, DownKey: enums.KeyDown}).Insert(
			&ebitencomponents.Velocity{
				X: 2,
				Y: 2,
				Z: 0,
			},
		)
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {

	}
}
