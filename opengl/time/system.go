package time

import "github.com/mrparano1d/ecs"

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			ecs.AddResource[*Resource](resourceMap, &Resource{delta: 0, lastCounter: 0})
		})
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {
		time := ecs.GetResource[*Resource](ctx.Resources)
		time.Update()
	}
}
