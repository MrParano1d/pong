package ebiten

import (
	"errors"
	"github.com/mrparano1d/pong/ebiten_plugin/enums"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

const (
	// TODO make screen width/height configurable
	screenWidth       = 512
	screenHeight      = 256
	defaultSaturation = 1.0
	defaultBrightness = 1.0
	defaultSkewX      = 0.0
	defaultSkewY      = 0.0
	defaultScaleX     = 1.0
	defaultScaleY     = 1.0
)

type renderCallback = func(surface interfaces.Surface) error

type updateCallback = func() error

// static check that we implement our renderer interface
var _ interfaces.Renderer = &Renderer{}

// Renderer is an implementation of a renderer
type Renderer struct {
	updateCallback
	renderCallback
	lastRenderError error
}

// Update calls the game's logical update function (the `Advance` method)
func (r *Renderer) Update() error {
	if r.updateCallback == nil {
		return errors.New("no update callback defined for ebiten renderer")
	}

	return r.updateCallback()
}

const drawError = "no render callback defined for ebiten renderer"

// Draw updates the screen with the given *ebiten.Image
func (r *Renderer) Draw(screen *ebiten.Image) {
	r.lastRenderError = nil

	if r.renderCallback == nil {
		r.lastRenderError = errors.New(drawError)
		return
	}

	r.lastRenderError = r.renderCallback(NewEbitenSurface(screen))
}

// Layout returns the renderer screen width and height
func (r *Renderer) Layout(_, _ int) (width, height int) {
	return screenWidth, screenHeight
}

// CreateRenderer creates an ebiten renderer instance
func CreateRenderer() (*Renderer, error) {
	result := &Renderer{}

	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	// ebiten.SetFullscreen(config.FullScreen)
	// ebiten.SetRunnableOnUnfocused(config.RunInBackground)
	// ebiten.SetVsyncEnabled(config.VsyncEnabled)
	// ebiten.SetMaxTPS(config.TicksPerSecond)

	return result, nil
}

// GetRendererName returns the name of the renderer
func (*Renderer) GetRendererName() string {
	return "Ebiten"
}

// SetWindowIcon sets the icon for the window, visible in the chrome of the window
func (*Renderer) SetWindowIcon(fileName string) {
	_, iconImage, err := ebitenutil.NewImageFromFile(fileName)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
}

// IsDrawingSkipped returns a bool for whether or not the drawing has been skipped
func (r *Renderer) IsDrawingSkipped() bool {
	return r.lastRenderError != nil
}

// Run initializes the renderer
func (r *Renderer) Run(f renderCallback, u updateCallback, width, height int, title string) error {
	r.renderCallback = f
	r.updateCallback = u

	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowSize(width, height)

	return ebiten.RunGame(r)
}

// CreateSurface creates a renderer surface from an existing surface
func (r *Renderer) CreateSurface(surface interfaces.Surface) (interfaces.Surface, error) {
	img := surface.(*EbitenSurface).Image

	sfcState := surfaceState{
		filter:     ebiten.FilterNearest,
		effect:     enums.DrawEffectNone,
		saturation: defaultSaturation,
		brightness: defaultBrightness,
		skewX:      defaultSkewX,
		skewY:      defaultSkewY,
		scaleX:     defaultScaleX,
		scaleY:     defaultScaleY,
		font:       GetDefaultFont(),
	}
	result := NewEbitenSurface(img, sfcState)

	return result, nil
}

func GetDefaultFont() font.Face {
	f, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("Parse Default Font: %v", err)
	}
	face, err := opentype.NewFace(
		f, &opentype.FaceOptions{
			Size:    16,
			DPI:     72,
			Hinting: font.HintingNone,
		},
	)
	if err != nil {
		log.Fatalf("New Default Face: %v", err)
	}
	return face
}

// NewSurface creates a new surface
func (r *Renderer) NewSurface(width, height int) interfaces.Surface {
	img := ebiten.NewImage(width, height)

	return NewEbitenSurface(img)
}

// IsFullScreen returns a boolean for whether or not the renderer is currently set to fullscreen
func (r *Renderer) IsFullScreen() bool {
	return ebiten.IsFullscreen()
}

// SetFullScreen sets the renderer to fullscreen, given a boolean
func (r *Renderer) SetFullScreen(fullScreen bool) {
	ebiten.SetFullscreen(fullScreen)
}

// SetVSyncEnabled enables vsync, given a boolean
func (r *Renderer) SetVSyncEnabled(vsync bool) {
	ebiten.SetVsyncEnabled(vsync)
}

// GetVSyncEnabled returns a boolean for whether or not vsync is enabled
func (r *Renderer) GetVSyncEnabled() bool {
	return ebiten.IsVsyncEnabled()
}

// GetCursorPos returns the current cursor position x,y coordinates
func (r *Renderer) GetCursorPos() (x, y int) {
	return ebiten.CursorPosition()
}

// CurrentFPS returns the current frames per second of the renderer
func (r *Renderer) CurrentFPS() float64 {
	return ebiten.CurrentFPS()
}

// CurrentTPS returns the current ticks per second of the renderer
func (r *Renderer) CurrentTPS() float64 {
	return ebiten.CurrentTPS()
}
