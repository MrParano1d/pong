package main

/*
Adapted from this tutorial: http://www.learnopengl.com/#!Getting-started/Textures
Shows how to use basic textures in openGL
*/

import (
	"flag"
	"fmt"
	"github.com/mrparano1d/ecs"
	"github.com/mrparano1d/ecs/core"
	"github.com/mrparano1d/pong/ebiten_plugin"
	"github.com/mrparano1d/pong/game"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpupprof", "", "write cpu profile to file")
var memprofile = flag.String("mempprof", "", "write memory profile to `file`")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}

		if err := pprof.StartCPUProfile(f); err != nil {
			panic(fmt.Errorf("failed to profile cpu: %v", err))
		}
		defer pprof.StopCPUProfile()
	}

	ebiten := ebiten_plugin.NewPlugin()

	app := ecs.NewApp()

	app.AddPlugin(core.NewPlugin(core.EnvDebug))
	app.AddPlugin(ebiten)
	app.AddPlugin(game.NewPlugin())

	if err := ebiten.Run(); err != nil {
		log.Fatal(err)
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer func() {
			// error handling omitted for example
			if err := f.Close(); err != nil {
				log.Printf("failed to close memprofile handler: %v\n", err)
			}
		}()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
