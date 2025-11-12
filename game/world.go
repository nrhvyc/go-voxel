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

	chunks := []Chunk{
		NewChunk(rl.NewVector3(0, 0, 0)),
		NewChunk(rl.NewVector3(16, 0, 0)),
		NewChunk(rl.NewVector3(0, 0, 16)),
		NewChunk(rl.NewVector3(16, 0, 16)),
	}

	for _, chunk := range chunks {
		world.Chunks[chunk.ID] = &chunk
	}

	return world
}
