package main

import "github.com/nrhvyc/go-voxel/game"

func main() {
	engine, err := game.NewEngine()
	if err != nil {
		panic(err)
	}
	engine.Run()
}
