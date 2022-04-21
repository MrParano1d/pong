package assets

type Asset interface {
	Width() float32
	Height() float32
	Create() error
	Bind()
	Unbind()
	Draw()
}
