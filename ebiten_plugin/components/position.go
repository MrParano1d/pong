package components

import (
	"github.com/mrparano1d/ecs"
	"reflect"
)

var positionComponentType = reflect.TypeOf(&Position{})

type Position struct {
	X float64
	Y float64
	Z float64
}

func (p *Position) Type() ecs.ComponentType {
	return positionComponentType
}
