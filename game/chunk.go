package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Chunk represents a 16x16x16 section of voxels
type Chunk struct {
	Voxels   [16][16][16]*Voxel
	Position rl.Vector3
}

// Render a single chunk
func (c *Chunk) render() {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				if voxel := c.Voxels[x][y][z]; voxel != nil {
					position := rl.Vector3Add(c.Position, rl.NewVector3(float32(x), float32(y), float32(z)))
					color := rl.NewColor(uint8(x*16), uint8(y*16), uint8(z*16), 255)
					rl.DrawCube(position, 0.9, 0.9, 0.9, color)
				}
			}
		}
	}
}
