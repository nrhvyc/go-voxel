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

	var chunks []Chunk
	chunkGenRadius := 2

	for x := -1 * chunkGenRadius; x <= chunkGenRadius; x++ {
		for z := -1 * chunkGenRadius; z <= chunkGenRadius; z++ {
			chunks = append(chunks, NewChunk(16*x, 16*z))
		}
	}

	for _, chunk := range chunks {
		world.Chunks[chunk.ID] = &chunk
	}

	return world
}
