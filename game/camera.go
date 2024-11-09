package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Create a new camera
func NewCamera() *Camera {
	return &Camera{
		Camera3D: rl.Camera3D{
			// Position: rl.NewVector3(0, float32(chunkHeight/2)+5, 20),
			Position: rl.NewVector3(0, 0, 0),
			Target:   rl.NewVector3(0, 0, 0),
			// Target:     rl.NewVector3(0, float32(chunkHeight/2), 0),
			Up:         rl.NewVector3(0, 1, 0),
			Fovy:       45.0,
			Projection: rl.CameraPerspective,
		},
	}
}
