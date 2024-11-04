package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	worldUp      = rl.Vector3{X: 0, Y: 1, Z: 0}
	worldForward = rl.Vector3{X: 0, Y: 0, Z: 1}
)

// World contains all chunks
type World struct {
	Chunks map[string]*Chunk
}

// Create a new world
func NewWorld() *World {
	world := &World{
		Chunks: make(map[string]*Chunk),
	}

	// Add a single chunk with some voxels
	chunk := &Chunk{
		Position: rl.NewVector3(0, 0, 0),
	}
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				if x == 0 || y == 0 || z == 0 || x == 15 || y == 15 || z == 15 {
					chunk.Voxels[x][y][z] = &Voxel{
						Position: rl.NewVector3(float32(x), float32(y), float32(z)),
						Type:     1,
					}
				}
			}
		}
	}
	world.Chunks["0,0,0"] = chunk

	return world
}
