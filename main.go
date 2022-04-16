package main

/*
Adapted from this tutorial: http://www.learnopengl.com/#!Getting-started/Textures
Shows how to use basic textures in openGL
*/

import (
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/pong/fps"
	"github.com/mrparano1d/pong/opengl"
	"log"
)

func main() {

	app := ecs.NewApp()

	app.AddPlugin(opengl.NewPlugin())
	app.AddPlugin(fps.NewPlugin())

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
