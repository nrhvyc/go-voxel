package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	chunkLength uint8 = 32  // x and z coords
	chunkHeight uint8 = 255 // y coords
)

type ChunkID string

// Chunk represents a 16x16x16 section of voxels
type Chunk struct {
	Voxels   [chunkLength][chunkHeight][chunkLength]*Voxel
	Position rl.Vector3
}

// Render a single chunk
func (c *Chunk) render() {
	for x := uint8(0); x < chunkLength; x++ {
		for y := uint8(0); y < chunkHeight; y++ {
			for z := uint8(0); z < chunkLength; z++ {
				if voxel := c.Voxels[x][y][z]; voxel != nil {
					position := rl.Vector3Add(c.Position, rl.NewVector3(float32(x), float32(y), float32(z)))
					color := rl.NewColor(uint8(x*16), uint8(y*16), uint8(z*16), 255)
					rl.DrawCube(position, 0.9, 0.9, 0.9, color)
				}
			}
		}
	}
}
