package components

import (
	"github.com/mrparano1d/ecs"
	"reflect"
)

var ballComponentType = reflect.TypeOf(&Ball{})

type Ball struct {
}

func (b *Ball) Type() ecs.ComponentType {
	return ballComponentType
}
