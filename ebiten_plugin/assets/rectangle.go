package assets

import (
	"github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"
	"image/color"
)

type Rectangle struct {
	Width  float64
	Height float64
	X      float64
	Y      float64
	Color  color.Color
	Border color.Color
}

var _ Handle = &Rectangle{}

func NewRectangle(width, height, x, y float64, color color.Color, border color.Color) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
		X:      x,
		Y:      y,
		Color:  color,
		Border: border,
	}
}

func (r *Rectangle) Translate(x, y float64) {
	r.X = x
	r.Y = y
}

func (r *Rectangle) Draw(surface interfaces.Surface) {
	surface.PushTranslation(r.X, r.Y)
	surface.DrawRect(r.Width, r.Height, r.Color, r.Border)
	surface.Pop()
}
