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
	Chunks map[ChunkID]*Chunk
}

// Create a new world
func NewWorld() *World {
	world := &World{
		Chunks: make(map[ChunkID]*Chunk),
	}

	// Add a single chunk with some voxels
	chunk := &Chunk{
		Position: rl.NewVector3(0, 0, 0),
	}

	for x := uint8(0); x < chunkLength; x++ {
		for y := uint8(chunkHeight / 2); y < chunkHeight/2+1; y++ {
			for z := uint8(0); z < chunkLength; z++ {
				// if x == 0 || y == 0 || z == 0 || x == 15 || y == 15 || z == 15 {
				height := Noise(x, y, z)
				chunk.Voxels[x][uint8(height)][z] = &Voxel{
					// Position: rl.NewVector3(float32(x), float32(y), float32(z)),
					Position: rl.NewVector3(float32(x), height, float32(z)),
					Type:     1,
				}
				// }
			}
		}
	}

	// y := chunkHeight / 2 // height
	// for x := uint8(0); x < chunkLength; x++ {
	// 	for z := uint8(0); z < chunkLength; z++ {
	// 		chunk.Voxels[x][y][z] = &Voxel{
	// 			Position: rl.NewVector3(float32(x), float32(y), float32(z)),
	// 			Type:     1,
	// 		}
	// 	}
	// }

	world.Chunks["0,0,0"] = chunk

	return world
}
