package ebiten

import (
	enums2 "github.com/mrparano1d/pong/ebiten_plugin/enums"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"
	"golang.org/x/image/font"
)

type colorMCacheKey uint32

const (
	maxAlpha       = 0xff
	cacheLimit     = 512
	transparency25 = 0.25
	transparency50 = 0.50
	transparency75 = 0.75
)

// static check that we implement the surface interface
var _ interfaces.Surface = &EbitenSurface{}

type colorMCacheEntry struct {
	colorMatrix ebiten.ColorM
	atime       int64
}
type EbitenSurface struct {
	Image          *ebiten.Image
	stateStack     []surfaceState
	stateCurrent   surfaceState
	monotonicClock int64
	colorMCache    map[colorMCacheKey]*colorMCacheEntry
}

func NewEbitenSurface(image *ebiten.Image, currentState ...surfaceState) *EbitenSurface {
	state := surfaceState{
		effect:     enums2.DrawEffectNone,
		saturation: defaultSaturation,
		brightness: defaultBrightness,
		skewX:      defaultSkewX,
		skewY:      defaultSkewY,
		scaleX:     defaultScaleX,
		scaleY:     defaultScaleY,
		font:       GetDefaultFont(),
	}
	if len(currentState) > 0 {
		state = currentState[0]
	}
	return &EbitenSurface{
		Image:        image,
		stateCurrent: state,
	}
}

// PushFont pushes a new font face to the state stack
func (s *EbitenSurface) PushFont(fnt font.Face) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.font = fnt
}

// PushTranslation pushes an x,y translation to the state stack
func (s *EbitenSurface) PushTranslation(x, y float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.x += x
	s.stateCurrent.y += y
}

// PushSkew pushes a skew to the state stack
func (s *EbitenSurface) PushSkew(skewX, skewY float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.skewX = skewX
	s.stateCurrent.skewY = skewY
}

// PushScale pushes a scale to the state stack
func (s *EbitenSurface) PushScale(scaleX, scaleY float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.scaleX = scaleX
	s.stateCurrent.scaleY = scaleY
}

// PushEffect pushes an effect to the state stack
func (s *EbitenSurface) PushEffect(effect enums2.DrawEffect) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.effect = effect
}

// PushFilter pushes a filter to the state stack
func (s *EbitenSurface) PushFilter(filter enums2.Filter) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.filter = ToEbitenFilter(filter)
}

// PushColor pushes a color to the stat stack
func (s *EbitenSurface) PushColor(c color.Color) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.color = c
}

// PushBrightness pushes a brightness value to the state stack
func (s *EbitenSurface) PushBrightness(brightness float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.brightness = brightness
}

// PushSaturation pushes a saturation value to the state stack
func (s *EbitenSurface) PushSaturation(saturation float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.saturation = saturation
}

// Pop pops a state off of the state stack
func (s *EbitenSurface) Pop() {
	count := len(s.stateStack)
	if count == 0 {
		panic("empty stack")
	}

	s.stateCurrent = s.stateStack[count-1]
	s.stateStack = s.stateStack[:count-1]
}

// PopN pops n states off the the state stack
func (s *EbitenSurface) PopN(n int) {
	for i := 0; i < n; i++ {
		s.Pop()
	}
}

// Render renders the given surface
func (s *EbitenSurface) Render(sfc interfaces.Surface) {
	opts := s.createDrawImageOptions()

	if s.stateCurrent.brightness != 1 || s.stateCurrent.saturation != 1 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	s.handleStateEffect(opts)

	s.Image.DrawImage(sfc.(*EbitenSurface).Image, opts)
}

// Renders the section of the surface, given the bounds
func (s *EbitenSurface) RenderSection(sfc interfaces.Surface, bound image.Rectangle) {
	opts := s.createDrawImageOptions()

	if s.stateCurrent.brightness != 0 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	s.handleStateEffect(opts)

	s.Image.DrawImage(sfc.(*EbitenSurface).Image.SubImage(bound).(*ebiten.Image), opts)
}

func (s *EbitenSurface) createDrawImageOptions() *ebiten.DrawImageOptions {
	opts := &ebiten.DrawImageOptions{}

	if s.stateCurrent.skewX != 0 || s.stateCurrent.skewY != 0 {
		opts.GeoM.Skew(s.stateCurrent.skewX, s.stateCurrent.skewY)
	}

	if s.stateCurrent.scaleX != 1.0 || s.stateCurrent.scaleY != 1.0 {
		opts.GeoM.Scale(s.stateCurrent.scaleX, s.stateCurrent.scaleY)
	}

	opts.GeoM.Translate(s.stateCurrent.x, s.stateCurrent.y)

	opts.Filter = s.stateCurrent.filter

	if s.stateCurrent.color != nil {
		opts.ColorM = s.colorToColorM(s.stateCurrent.color)
	}

	return opts
}

func (s *EbitenSurface) handleStateEffect(opts *ebiten.DrawImageOptions) {
	switch s.stateCurrent.effect {
	case enums2.DrawEffectPctTransparency25:
		opts.ColorM.Translate(0, 0, 0, -transparency25)
	case enums2.DrawEffectPctTransparency50:
		opts.ColorM.Translate(0, 0, 0, -transparency50)
	case enums2.DrawEffectPctTransparency75:
		opts.ColorM.Translate(0, 0, 0, -transparency75)
	case enums2.DrawEffectModulate:
		opts.CompositeMode = ebiten.CompositeModeLighter
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/822
	case enums2.DrawEffectBurn:
	case enums2.DrawEffectNormal:
	case enums2.DrawEffectMod2XTrans:
	case enums2.DrawEffectMod2X:
	case enums2.DrawEffectNone:
		opts.CompositeMode = ebiten.CompositeModeSourceOver
	}
}

// DrawTextf renders the string to the surface with the given format string and a set of parameters
func (s *EbitenSurface) DrawText(txt string, clr color.Color) {
	if s.stateCurrent.font == nil {
		log.Fatal("ebitenSurface: No font face registered for the current surface state")
	}

	// text.Draw(s.Image, txt, s.stateCurrent.font, int(s.stateCurrent.x), int(s.stateCurrent.y+16), clr)
	ebitenutil.DebugPrintAt(s.Image, txt, int(s.stateCurrent.x), int(s.stateCurrent.y))
}

// DrawLine draws a line
func (s *EbitenSurface) DrawLine(x, y float64, fillColor color.Color) {
	ebitenutil.DrawLine(
		s.Image,
		s.stateCurrent.x,
		s.stateCurrent.y,
		s.stateCurrent.x+x,
		s.stateCurrent.y+y,
		fillColor,
	)
}

// DrawRect draws a rectangle
func (s *EbitenSurface) DrawRect(width, height float64, fillColor color.Color, borderColor color.Color) {

	if borderColor != nil {
		// draw rect border
		for _, l := range borderRect(s.stateCurrent.x, s.stateCurrent.y, width, height) {
			ebitenutil.DrawLine(s.Image, l.X1, l.Y1, l.X2, l.Y2, borderColor)
		}
	}

	_, _, _, a := fillColor.RGBA()

	if a != 0 {
		ebitenutil.DrawRect(
			s.Image,
			s.stateCurrent.x,
			s.stateCurrent.y,
			width,
			height,
			fillColor,
		)
	}
}

// Clear clears the entire surface, filling with the given color
func (s *EbitenSurface) Clear(fillColor color.Color) {
	s.Image.Fill(fillColor)
}

// GetSize gets the size of the surface
func (s *EbitenSurface) GetSize() (x, y int) {
	return s.Image.Size()
}

// GetDepth returns the depth of this surface in the stack
func (s *EbitenSurface) GetDepth() int {
	return len(s.stateStack)
}

// ReplacePixels replaces pixels in the surface with the given pixels
func (s *EbitenSurface) ReplacePixels(pixels []byte) {
	s.Image.ReplacePixels(pixels)
}

// Screenshot returns an *image.RGBA of the surface
func (s *EbitenSurface) Screenshot() *image.RGBA {
	width, height := s.GetSize()
	bounds := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: width, Y: height}}
	rgba := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgba.Set(x, y, s.Image.At(x, y))
		}
	}

	return rgba
}

func (s *EbitenSurface) now() int64 {
	s.monotonicClock++
	return s.monotonicClock
}

// colorToColorM converts a normal color to a color matrix
func (s *EbitenSurface) colorToColorM(clr color.Color) ebiten.ColorM {
	// RGBA() is in [0 - 0xffff]. Adjust them in [0 - 0xff].
	cr, cg, cb, ca := clr.RGBA()
	cr >>= 8
	cg >>= 8
	cb >>= 8
	ca >>= 8

	if ca == 0 {
		emptyColorM := ebiten.ColorM{}
		emptyColorM.Scale(0, 0, 0, 0)

		return emptyColorM
	}

	// nolint:gomnd // byte values
	key := colorMCacheKey(cr | (cg << 8) | (cb << 16) | (ca << 24))
	e, ok := s.colorMCache[key]

	if ok {
		e.atime = s.now()
		return e.colorMatrix
	}

	if len(s.colorMCache) > cacheLimit {
		oldest := int64(math.MaxInt64)
		oldestKey := colorMCacheKey(0)

		for key, c := range s.colorMCache {
			if c.atime < oldest {
				oldestKey = key
				oldest = c.atime
			}
		}

		delete(s.colorMCache, oldestKey)
	}

	cm := ebiten.ColorM{}
	rf := float64(cr) / float64(ca)
	gf := float64(cg) / float64(ca)
	bf := float64(cb) / float64(ca)
	af := float64(ca) / maxAlpha
	cm.Scale(rf, gf, bf, af)

	e = &colorMCacheEntry{
		colorMatrix: cm,
		atime:       s.now(),
	}

	s.colorMCache[key] = e

	return e.colorMatrix
}
