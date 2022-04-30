package assets

import (
	"github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"
	"image/color"
)

type Text struct {
	X     float64
	Y     float64
	Color color.Color
	Text  string
}

var _ Handle = &Text{}

func NewText(text string, x, y float64, color color.Color) *Text {
	return &Text{
		X:     x,
		Y:     y,
		Color: color,
		Text:  text,
	}
}

func (t *Text) Width() float64 {
	return 0
}

func (t *Text) Height() float64 {
	return 0
}

func (t *Text) Translate(x, y float64) {
	t.X = x
	t.Y = y
}

func (t *Text) Draw(surface interfaces.Surface) {
	surface.PushTranslation(t.X, t.Y)
	surface.DrawText(t.Text, t.Color)
	surface.Pop()
}
