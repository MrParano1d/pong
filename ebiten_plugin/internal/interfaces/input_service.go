package interfaces

import (
	enums2 "github.com/mrparano1d/pong/ebiten_plugin/enums"
)

type InputService interface {
	CursorPosition() (x, y int)
	InputChars() []rune

	IsKeyPressed(key enums2.Key) bool

	IsKeyJustPressed(key enums2.Key) bool

	IsKeyJustReleased(key enums2.Key) bool

	IsMouseButtonPressed(button enums2.MouseButton) bool

	IsMouseButtonJustPressed(button enums2.MouseButton) bool
	IsMouseButtonJustReleased(button enums2.MouseButton) bool

	KeyPressDuration(key enums2.Key) int
}
