package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Voxel represents a single cube in the world
type Voxel struct {
	Position rl.Vector3
	Type     int
}
