package assets

import "github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"

type Handle interface {
	Height() float64
	Width() float64
	Translate(x, y float64)
	Draw(surface interfaces.Surface)
}
