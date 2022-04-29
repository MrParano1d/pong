package camera

import (
	"github.com/mrparano1d/ecs"
)

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(
			func(resourceMap ecs.ResourceMap) {

			},
		)
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {

	}
}
