package assets

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"github.com/mrparano1d/pong/opengl/camera"
	"github.com/mrparano1d/pong/opengl/shapes"
	"go.uber.org/zap"
	"image/color"
	time2 "time"
)

func Setup() ecs.StartUpSystem {
	return func(commands ecs.Commands) {
		commands.InvokeResource(func(resourceMap ecs.ResourceMap) {
			logger := ecs.GetResource[*zap.Logger](resourceMap)
			triangleAsset := shapes.NewTriangle(
				[3]float32{-0.5, -0.5, 0.0},
				[3]float32{0.5, -0.5, 0.0},
				[3]float32{0.0, 0.5, 0.0},
				color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			)
			if err := triangleAsset.Create(); err != nil {
				logger.Fatal("failed to load triangle asset", zap.Error(err))
			}
			ecs.AddResource[*shapes.Triangle](resourceMap, &triangleAsset)

			model := mgl32.Scale3D(1, 1, 1.0)
			//modelScale := mgl32.Scale3D(1.2, 1.2, 1.2)
			//model = model.Mul4(modelScale)

			ecs.AddResource[mgl32.Mat4](resourceMap, model)
		})
	}
}

func System() ecs.System {
	return func(ctx ecs.SystemContext) {
		time2.Sleep(1 * time2.Millisecond)
		model := ecs.GetResource[mgl32.Mat4](ctx.Resources)
		asset := ecs.GetResource[*shapes.Triangle](ctx.Resources)
		cam := ecs.GetResource[*camera.Camera](ctx.Resources)
		time := ecs.GetResource[*core.Time](ctx.Resources)

		rotate := mgl32.HomogRotate3DY(20.0 * float32(time.Delta()))
		model = model.Mul4(rotate)

		modelViewProj := cam.ViewProjection().Mul4(model)
		cam.Update()

		asset.Draw(modelViewProj)
	}
}
