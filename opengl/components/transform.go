package components

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/ecs"
	"reflect"
)

var transformType = reflect.TypeOf(&Transform2D{})

type Transform2D struct {
	Scale    mgl32.Vec2
	Rotation float32
}

func (t *Transform2D) Type() ecs.ComponentType {
	return transformType
}
