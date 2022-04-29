package assets

import "github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"

type Handle interface {
	Translate(x, y float64)
	Draw(surface interfaces.Surface)
}
