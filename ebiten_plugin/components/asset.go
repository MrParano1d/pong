package components

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/assets"
	"reflect"
)

var assetComponentType = reflect.TypeOf(&Asset{})

type Asset struct {
	Handle assets.Handle
}

var _ ecs.Component = &Asset{}

func (a *Asset) Type() ecs.ComponentType {
	return assetComponentType
}
