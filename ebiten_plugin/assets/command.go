package assets

import (
	"github.com/mrparano1d/ecs"
	"sync"
)

type CommandAssets map[string]Handle

type Command struct {
	mutex  *sync.Mutex
	assets CommandAssets
}

var _ ecs.Command = &Command{}

func NewCommand(assets CommandAssets) *Command {
	if assets == nil {
		assets = map[string]Handle{}
	}
	return &Command{
		mutex:  &sync.Mutex{},
		assets: assets,
	}
}

func (c *Command) Add(name string, handle Handle) *Command {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.assets[name] = handle
	return c
}

func (c *Command) Write(world *ecs.World) {
	manager := ecs.GetResource[*Manager](world.Resources())
	for name, handle := range c.assets {
		manager.Add(name, handle)
	}
}
