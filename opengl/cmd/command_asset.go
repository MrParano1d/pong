package cmd

import (
	"fmt"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/components"
	"go.uber.org/zap"
)

type AssetCommand struct {
}

var _ ecs.Command = &AssetCommand{}

func (a *AssetCommand) Write(world *ecs.World) {
	query := ecs.NewQuery(world)
	logger := ecs.GetResource[*zap.Logger](world.Resources())

	for _, eID := range query.Find(ecs.NewFilter(ecs.WithComponentFilter(&components.Asset{}))) {
		asset := ecs.GetComponent[*components.Asset](world.Entities(), eID)
		if err := asset.Handle.Create(); err != nil {
			logger.Error("failed to create asset", zap.String("type", fmt.Sprintf("%T", asset)))
		}
	}
}
