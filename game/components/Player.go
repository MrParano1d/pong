package components

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/ebiten_plugin/enums"
	"reflect"
)

var playerComponentType = reflect.TypeOf(&Player{})

type Player struct {
	UpKey   enums.Key
	DownKey enums.Key
}

func (p *Player) Type() ecs.ComponentType {
	return playerComponentType
}
