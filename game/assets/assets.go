package assets

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/assets"
	ebitencomponents "github.com/mrparano1d/pong/ebiten_plugin/components"
	"github.com/mrparano1d/pong/ebiten_plugin/enums"
	"github.com/mrparano1d/pong/game/components"
	"image/color"
)

const (
	PlayerHeight    = 28
	PlayerWidth     = 4
	BallWidth       = 6
	BallHeight      = 6
	PlayerPositionY = (256 / 2) - (PlayerHeight / 2)
	BallPositionX   = (512 / 2) - (BallWidth / 2)
	BallPositionY   = (256 / 2) - (BallHeight / 2)
)

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {

		// Assets
		player1 := assets.NewRectangle(PlayerWidth, PlayerHeight, 0, 0, color.RGBA{255, 255, 255, 255}, nil)
		player2 := assets.NewRectangle(PlayerWidth, PlayerHeight, 0, 0, color.RGBA{255, 255, 255, 255}, nil)

		middleline := assets.NewRectangle(2, 256, 0, 0, color.RGBA{255, 255, 255, 255}, nil)

		ball := assets.NewRectangle(BallWidth, BallHeight, 0, 0, color.RGBA{255, 255, 255, 255}, nil)

		score1 := assets.NewText("0", 0, 0, color.RGBA{255, 255, 255, 255})
		score2 := assets.NewText("0", 0, 0, color.RGBA{255, 255, 255, 255})

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
			).Add(
				"ball",
				ball,
			).Add(
				"score1",
				score1,
			).Add(
				"score2",
				score2,
			),
		)

		// Score1
		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: score1}).Insert(
			&ebitencomponents.Position{
				X: 20,
				Y: 10,
			},
		)

		// Score2
		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: score2}).Insert(
			&ebitencomponents.Position{
				X: 512 - 30,
				Y: 10,
			},
		)

		// Middleline
		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: middleline}).Insert(
			&ebitencomponents.Position{
				X: (512 / 2) - (middleline.Width() / 2),
				Y: 0,
			},
		)

		// Ball
		commands.Spawn().Insert(
			&ebitencomponents.Asset{
				Handle: ball,
			},
		).Insert(
			&ebitencomponents.Position{
				X: BallPositionX,
				Y: BallPositionY,
			},
		).Insert(
			&ebitencomponents.Velocity{
				X: 2,
				Y: 2,
				Z: 0,
			},
		).Insert(
			&ebitencomponents.WorldCollider{
				Bounce: true,
			},
		).Insert(
			&ebitencomponents.Collision{
				Width:  ball.Width(),
				Height: ball.Height(),
			},
		).Insert(&components.Ball{})

		// Player 1
		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: player1}).Insert(
			&ebitencomponents.Position{
				X: 14,
				Y: PlayerPositionY,
			},
		).Insert(&components.Player{UpKey: enums.KeyW, DownKey: enums.KeyS}).Insert(
			&ebitencomponents.Velocity{
				X: 2,
				Y: 2,
				Z: 0,
			},
		).Insert(
			&ebitencomponents.Collision{
				Width:  player1.Width(),
				Height: player1.Height(),
			},
		).Insert(
			&ebitencomponents.WorldCollider{
				Bounce: false,
			},
		)

		// Player 2
		commands.Spawn().Insert(&ebitencomponents.Asset{Handle: player2}).Insert(
			&ebitencomponents.Position{
				X: 512 - 14,
				Y: PlayerPositionY,
			},
		).Insert(&components.Player{UpKey: enums.KeyUp, DownKey: enums.KeyDown}).Insert(
			&ebitencomponents.Velocity{
				X: 2,
				Y: 2,
				Z: 0,
			},
		).Insert(
			&ebitencomponents.Collision{
				Width:  player2.Width(),
				Height: player2.Height(),
			},
		).Insert(
			&ebitencomponents.WorldCollider{
				Bounce: false,
			},
		)
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {

	}
}
