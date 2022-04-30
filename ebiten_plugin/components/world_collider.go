package components

import (
	"github.com/mrparano1d/ecs"
	"reflect"
)

var worldColliderComponentType = reflect.TypeOf(&WorldCollider{})

type WorldCollider struct {
	Bounce bool
}

func (w *WorldCollider) Type() ecs.ComponentType {
	return worldColliderComponentType
}
