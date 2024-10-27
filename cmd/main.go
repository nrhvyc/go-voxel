package main

import "github.com/nrhvyc/voxel-engine/voxel"

func main() {
	engine, err := voxel.NewEngine()
	if err != nil {
		panic(err)
	}
	engine.Run()
}
