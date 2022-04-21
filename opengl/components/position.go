package components

import (
	"github.com/mrparano1d/ecs"
	"reflect"
)

var positionType = reflect.TypeOf(&Position{})

type Position struct {
	X float32
	Y float32
	Z float32
}

func (p *Position) Type() ecs.ComponentType {
	return positionType
}
