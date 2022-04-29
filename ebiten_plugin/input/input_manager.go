package input

import "github.com/mrparano1d/pong/ebiten_plugin/internal/interfaces"

type InputManager struct {
	service interfaces.InputService
}

var _ interfaces.InputManager = &InputManager{}

func NewInputManager(inputService interfaces.InputService) *InputManager {
	return &InputManager{
		service: inputService,
	}
}

func (im *InputManager) SetInputService(service interfaces.InputService) {
	im.service = service
}
func (im *InputManager) GetInputService() interfaces.InputService {
	return im.service
}
