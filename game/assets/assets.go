package assets

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/opengl/camera"
	"github.com/mrparano1d/pong/opengl/shapes"
	"github.com/mrparano1d/pong/opengl/time"
	"github.com/mrparano1d/pong/opengl/window"
	"go.uber.org/zap"
	"image/color"
	"reflect"
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

			rectangleAsset := shapes.NewRectangle(0, 0, 200, 200,
				color.RGBA{
					R: 0,
					G: 255,
					B: 0,
					A: 255,
				},
			)
			if err := rectangleAsset.Create(); err != nil {
				logger.Fatal("failed to load rectangle asset", zap.Error(err))
			}
			ecs.AddResource[*shapes.Rectangle](resourceMap, &rectangleAsset)

			model := mgl32.Translate3D(0, 0, 0)
			//modelScale := mgl32.Scale3D(1.2, 1.2, 1.2)
			//model = model.Mul4(modelScale)

			ecs.AddResource[mgl32.Mat4](resourceMap, model)
		})
	}
}

func System() ecs.System {
	//var rotationTimer float32

	position := mgl32.Vec2{0, 0}
	velocity := mgl32.Vec2{150, 150}

	return func(ctx ecs.SystemContext) {
		model := ecs.GetResource[mgl32.Mat4](ctx.Resources)
		asset := ecs.GetResource[*shapes.Rectangle](ctx.Resources)
		cam := ecs.GetResource[*camera.Cameras](ctx.Resources)
		win := ecs.GetResource[*window.Resource](ctx.Resources)
		t := ecs.GetResource[*time.Resource](ctx.Resources)

		projections := ecs.GetResource[*camera.Projections](ctx.Resources)

		position = position.Add(velocity.Mul(float32(t.Delta())))

		fmt.Println("position", position)

		fmt.Println("asset window", win.Width, "x", win.Height, reflect.DeepEqual(cam.Orthographic.Projection(), projections.Orthographic))

		if position.X()+float32(asset.Width) >= win.Width || position.X() <= 0 {
			velocity[0] *= -1
		}
		if position.Y()+float32(asset.Height) >= win.Height || position.Y() <= 0 {
			velocity[1] *= -1
		}

		// position
		model = model.Mul4(mgl32.Translate3D(position.X(), position.Y(), 0.0))
		//// move origin of rotation to center of quad
		//model = model.Mul4(mgl32.Translate3D(0.5*size.X(), 0.5*size.Y(), 0.0))
		//// rotation
		//model = model.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(0)))
		//// move origin back
		//model = model.Mul4(mgl32.Translate3D(-0.5*size.X(), -0.5*size.Y(), 0.0))
		//// scale
		//model = model.Mul4(mgl32.Scale3D(size.X(), size.Y(), 1.0))

		//modelViewProj := cam.ViewProjection().Mul4(model)
		//cam.Update()

		asset.Draw(cam.Orthographic.ViewProjection().Mul4(model))

		//ecs.AddResource[mgl32.Mat4](ctx.Resources, model)
	}
}
