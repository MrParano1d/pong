package interfaces

type InputManager interface {
	SetInputService(service InputService)
	GetInputService() InputService
}
