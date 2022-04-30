package components

import (
	"github.com/mrparano1d/ecs"
	"reflect"
)

var collisionComponentType = reflect.TypeOf(&Collision{})

type Collision struct {
	Width  float64
	Height float64
}

func (c *Collision) Type() ecs.ComponentType {
	return collisionComponentType
}
