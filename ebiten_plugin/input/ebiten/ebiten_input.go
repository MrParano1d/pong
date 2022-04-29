// Package ebiten provides graphics and input API to develop a 2D game.
package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	enums2 "github.com/mrparano1d/pong/ebiten_plugin/enums"
	"github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"
)

var (
	//nolint:gochecknoglobals // This is a constant in all but by name, no constant map in go
	keyToEbiten = map[enums2.Key]ebiten.Key{
		enums2.Key0:            ebiten.Key0,
		enums2.Key1:            ebiten.Key1,
		enums2.Key2:            ebiten.Key2,
		enums2.Key3:            ebiten.Key3,
		enums2.Key4:            ebiten.Key4,
		enums2.Key5:            ebiten.Key5,
		enums2.Key6:            ebiten.Key6,
		enums2.Key7:            ebiten.Key7,
		enums2.Key8:            ebiten.Key8,
		enums2.Key9:            ebiten.Key9,
		enums2.KeyA:            ebiten.KeyA,
		enums2.KeyB:            ebiten.KeyB,
		enums2.KeyC:            ebiten.KeyC,
		enums2.KeyD:            ebiten.KeyD,
		enums2.KeyE:            ebiten.KeyE,
		enums2.KeyF:            ebiten.KeyF,
		enums2.KeyG:            ebiten.KeyG,
		enums2.KeyH:            ebiten.KeyH,
		enums2.KeyI:            ebiten.KeyI,
		enums2.KeyJ:            ebiten.KeyJ,
		enums2.KeyK:            ebiten.KeyK,
		enums2.KeyL:            ebiten.KeyL,
		enums2.KeyM:            ebiten.KeyM,
		enums2.KeyN:            ebiten.KeyN,
		enums2.KeyO:            ebiten.KeyO,
		enums2.KeyP:            ebiten.KeyP,
		enums2.KeyQ:            ebiten.KeyQ,
		enums2.KeyR:            ebiten.KeyR,
		enums2.KeyS:            ebiten.KeyS,
		enums2.KeyT:            ebiten.KeyT,
		enums2.KeyU:            ebiten.KeyU,
		enums2.KeyV:            ebiten.KeyV,
		enums2.KeyW:            ebiten.KeyW,
		enums2.KeyX:            ebiten.KeyX,
		enums2.KeyY:            ebiten.KeyY,
		enums2.KeyZ:            ebiten.KeyZ,
		enums2.KeyApostrophe:   ebiten.KeyApostrophe,
		enums2.KeyBackslash:    ebiten.KeyBackslash,
		enums2.KeyBackspace:    ebiten.KeyBackspace,
		enums2.KeyCapsLock:     ebiten.KeyCapsLock,
		enums2.KeyComma:        ebiten.KeyComma,
		enums2.KeyDelete:       ebiten.KeyDelete,
		enums2.KeyDown:         ebiten.KeyDown,
		enums2.KeyEnd:          ebiten.KeyEnd,
		enums2.KeyEnter:        ebiten.KeyEnter,
		enums2.KeyEqual:        ebiten.KeyEqual,
		enums2.KeyEscape:       ebiten.KeyEscape,
		enums2.KeyF1:           ebiten.KeyF1,
		enums2.KeyF2:           ebiten.KeyF2,
		enums2.KeyF3:           ebiten.KeyF3,
		enums2.KeyF4:           ebiten.KeyF4,
		enums2.KeyF5:           ebiten.KeyF5,
		enums2.KeyF6:           ebiten.KeyF6,
		enums2.KeyF7:           ebiten.KeyF7,
		enums2.KeyF8:           ebiten.KeyF8,
		enums2.KeyF9:           ebiten.KeyF9,
		enums2.KeyF10:          ebiten.KeyF10,
		enums2.KeyF11:          ebiten.KeyF11,
		enums2.KeyF12:          ebiten.KeyF12,
		enums2.KeyGraveAccent:  ebiten.KeyGraveAccent,
		enums2.KeyHome:         ebiten.KeyHome,
		enums2.KeyInsert:       ebiten.KeyInsert,
		enums2.KeyKP0:          ebiten.KeyKP0,
		enums2.KeyKP1:          ebiten.KeyKP1,
		enums2.KeyKP2:          ebiten.KeyKP2,
		enums2.KeyKP3:          ebiten.KeyKP3,
		enums2.KeyKP4:          ebiten.KeyKP4,
		enums2.KeyKP5:          ebiten.KeyKP5,
		enums2.KeyKP6:          ebiten.KeyKP6,
		enums2.KeyKP7:          ebiten.KeyKP7,
		enums2.KeyKP8:          ebiten.KeyKP8,
		enums2.KeyKP9:          ebiten.KeyKP9,
		enums2.KeyKPAdd:        ebiten.KeyKPAdd,
		enums2.KeyKPDecimal:    ebiten.KeyKPDecimal,
		enums2.KeyKPDivide:     ebiten.KeyKPDivide,
		enums2.KeyKPEnter:      ebiten.KeyKPEnter,
		enums2.KeyKPEqual:      ebiten.KeyKPEqual,
		enums2.KeyKPMultiply:   ebiten.KeyKPMultiply,
		enums2.KeyKPSubtract:   ebiten.KeyKPSubtract,
		enums2.KeyLeft:         ebiten.KeyLeft,
		enums2.KeyLeftBracket:  ebiten.KeyLeftBracket,
		enums2.KeyMenu:         ebiten.KeyMenu,
		enums2.KeyMinus:        ebiten.KeyMinus,
		enums2.KeyNumLock:      ebiten.KeyNumLock,
		enums2.KeyPageDown:     ebiten.KeyPageDown,
		enums2.KeyPageUp:       ebiten.KeyPageUp,
		enums2.KeyPause:        ebiten.KeyPause,
		enums2.KeyPeriod:       ebiten.KeyPeriod,
		enums2.KeyPrintScreen:  ebiten.KeyPrintScreen,
		enums2.KeyRight:        ebiten.KeyRight,
		enums2.KeyRightBracket: ebiten.KeyRightBracket,
		enums2.KeyScrollLock:   ebiten.KeyScrollLock,
		enums2.KeySemicolon:    ebiten.KeySemicolon,
		enums2.KeySlash:        ebiten.KeySlash,
		enums2.KeySpace:        ebiten.KeySpace,
		enums2.KeyTab:          ebiten.KeyTab,
		enums2.KeyUp:           ebiten.KeyUp,
		enums2.KeyAlt:          ebiten.KeyAlt,
		enums2.KeyControl:      ebiten.KeyControl,
		enums2.KeyShift:        ebiten.KeyShift,
	}
	//nolint:gochecknoglobals // This is a constant in all but by name, no constant map in go
	mouseButtonToEbiten = map[enums2.MouseButton]ebiten.MouseButton{
		enums2.MouseButtonLeft:   ebiten.MouseButtonLeft,
		enums2.MouseButtonMiddle: ebiten.MouseButtonMiddle,
		enums2.MouseButtonRight:  ebiten.MouseButtonRight,
	}
)

// InputService provides an abstraction on ebiten to support handling input events
type InputService struct{}

var _ interfaces.InputService = &InputService{}

// CursorPosition returns a position of a mouse cursor relative to the game screen (window).
func (is InputService) CursorPosition() (x, y int) {
	return ebiten.CursorPosition()
}

// InputChars return "printable" runes read from the keyboard at the time update is called.
func (is InputService) InputChars() []rune {
	return ebiten.InputChars()
}

// IsKeyPressed checks if the provided key is down.
func (is InputService) IsKeyPressed(key enums2.Key) bool {
	return ebiten.IsKeyPressed(keyToEbiten[key])
}

// IsKeyJustPressed checks if the provided key is just transitioned from up to down.
func (is InputService) IsKeyJustPressed(key enums2.Key) bool {
	return inpututil.IsKeyJustPressed(keyToEbiten[key])
}

// IsKeyJustReleased checks if the provided key is just transitioned from down to up.
func (is InputService) IsKeyJustReleased(key enums2.Key) bool {
	return inpututil.IsKeyJustReleased(keyToEbiten[key])
}

// IsMouseButtonPressed checks if the provided mouse button is down.
func (is InputService) IsMouseButtonPressed(button enums2.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustPressed checks if the provided mouse button is just transitioned from up to down.
func (is InputService) IsMouseButtonJustPressed(button enums2.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustReleased checks if the provided mouse button is just transitioned from down to up.
func (is InputService) IsMouseButtonJustReleased(button enums2.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(mouseButtonToEbiten[button])
}

// KeyPressDuration returns how long the key is pressed in frames.
func (is InputService) KeyPressDuration(key enums2.Key) int {
	return inpututil.KeyPressDuration(keyToEbiten[key])
}
