package ebiten

import (
	"github.com/mrparano1d/pong/ebiten_plugin/enums"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type surfaceState struct {
	x              float64
	y              float64
	filter         ebiten.Filter
	color          color.Color
	brightness     float64
	saturation     float64
	effect         enums.DrawEffect
	skewX, skewY   float64
	scaleX, scaleY float64
	font           font.Face
}
