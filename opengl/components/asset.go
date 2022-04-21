package components

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/assets"
	"reflect"
)

var assetType = reflect.TypeOf(&Asset{})

type Asset struct {
	Handle assets.Asset
}

func (a *Asset) Type() ecs.ComponentType {
	return assetType
}
