package interfaces

import (
	enums2 "github.com/mrparano1d/pong/ebiten_plugin/enums"
	"image"
	"image/color"

	"golang.org/x/image/font"
)

type Surface interface {
	Renderable

	Clear(color color.Color)
	DrawRect(width, height float64, fillColor color.Color, borderColor color.Color)
	DrawLine(x, y float64, color color.Color)
	DrawText(txt string, clr color.Color)
	GetSize() (width, height int)
	GetDepth() int
	Pop()
	PopN(n int)
	PushColor(color color.Color)
	PushEffect(effect enums2.DrawEffect)
	PushFilter(filter enums2.Filter)
	PushFont(fnt font.Face)
	PushTranslation(x, y float64)
	PushSkew(x, y float64)
	PushScale(x, y float64)
	PushBrightness(brightness float64)
	PushSaturation(saturation float64)
	// RenderSection renders a section of the surface enclosed by bounds
	RenderSection(surface Surface, bound image.Rectangle)
	ReplacePixels(pixels []byte)
	Screenshot() *image.RGBA
}
