package game

import (
	"fmt"
	"image/color"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	chunkLength uint8 = 16 // x and z coords
	chunkHeight uint8 = 16 // y coords
)

var (
	chunkBoundingBoxColor = color.RGBA{255, 0, 0, 255}
)

type ChunkID string

func (id ChunkID) String() string {
	return string(id)
}

// Chunk represents a 16x16x16 section of voxels
type Chunk struct {
	Voxels [chunkLength][chunkHeight][chunkLength]*Voxel

	ID            ChunkID
	worldPosition rl.Vector3
	boundingBox   rl.BoundingBox

	debugColor color.RGBA

	// offsets so worldPosition is the center of a chunk when rendered
	xRenderOffset, zRenderOffset uint8
}

func NewChunk(xPos, zPos int) Chunk {
	chunk := Chunk{
		worldPosition: rl.NewVector3(float32(xPos), 0, float32(zPos)),
		ID: ChunkID(fmt.Sprintf("%d,%d",
			xPos,
			zPos,
		)),
	}

	offset := uint8(chunkLength / 2)

	chunk.xRenderOffset = uint8(offset)
	chunk.zRenderOffset = uint8(offset)

	// Generate the Voxels for the chunk
	for x := range chunkLength {
		for y := range uint8(chunkHeight/2 + 1) {
			for z := range chunkLength {
				height := Noise(x, y, z)
				height = 0

				chunk.Voxels[x][uint8(height)][z] = &Voxel{
					// NOTE: chunk is only single layer deep right now
					// Position: rl.NewVector3(float32(x), float32(y), float32(z)),
					Position: rl.NewVector3(
						float32(x),
						height,
						float32(z),
					),
					Type: 1,
				}
			}
		}
	}

	// Chunks don't move, so we only need to calculate the boundingBox once
	chunk.boundingBox = rl.BoundingBox{
		Min: rl.Vector3{
			X: chunk.worldPosition.X - 0.5,
			Y: 0 - 0.5,
			Z: chunk.worldPosition.Z - 0.5,
		},
		Max: rl.Vector3{
			X: chunk.worldPosition.X + float32(chunkLength) - 0.5,
			Y: 255,
			Z: chunk.worldPosition.Z + float32(chunkLength) - 0.5,
		},
	}
	// gonna use rl.DrawCubeWires to draw the bounding box

	// Basically the color that'll render for voxels in the chunk
	chunk.debugColor = rl.NewColor(
		rand.N(uint8(255)),
		rand.N(uint8(255)),
		rand.N(uint8(255)),
		255)

	return chunk
}

// Render a single chunk
func (c *Chunk) render() {

	// Eventually add a check for whether the chunk is in view of the frustum
	rl.DrawBoundingBox(c.boundingBox, chunkBoundingBoxColor)

	for x := range chunkLength {
		for y := range chunkHeight {
			for z := range chunkLength {

				// Skip missing voxels
				if c.Voxels[x][y][z] == nil {
					continue
				}

				// Draw voxels in the chunk with the c.worldPosition
				// being at the center
				position := rl.Vector3Add(
					c.worldPosition,
					rl.NewVector3(
						float32(x),
						float32(y),
						float32(z),
					),
				)

				// used for individual random voxel color
				// color := rl.NewColor(uint8(x*16), uint8(y*16), uint8(z*16), 255)

				rl.DrawCube(position, 1, 1, 1, c.debugColor)
				rl.DrawCubeWires(position, 1, 1, 1, VoxelOutlineColor)
			}
		}
	}
}
