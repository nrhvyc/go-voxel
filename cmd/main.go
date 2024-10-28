package main

import "github.com/nrhvyc/go-voxel/voxel"

func main() {
	engine, err := voxel.NewEngine()
	if err != nil {
		panic(err)
	}
	engine.Run()
}
