package cmd

import "github.com/mrparano1d/ecs"

type ResizeCallback func(world *ecs.World)

type WindowCommand struct {
	ResizeCallback ResizeCallback
}

var _ ecs.Command = &WindowCommand{}

func (r *WindowCommand) Write(world *ecs.World) {
	r.ResizeCallback(world)
}
