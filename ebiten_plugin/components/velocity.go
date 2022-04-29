package components

import (
	"github.com/mrparano1d/ecs"
	"reflect"
)

var velocityComponentType = reflect.TypeOf(&Velocity{})

type Velocity struct {
	X float64
	Y float64
	Z float64
}

func (v *Velocity) Type() ecs.ComponentType {
	return velocityComponentType
}
