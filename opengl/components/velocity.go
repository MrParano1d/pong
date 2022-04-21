package components

import (
	"github.com/mrparano1d/ecs"
	"reflect"
)

var velocityType = reflect.TypeOf(&Velocity{})

type Velocity struct {
	X float32
	Y float32
	Z float32
}

func (t *Velocity) Type() ecs.ComponentType {
	return velocityType
}
