package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ninetyDegRadians = 90.0 * (math.Pi / 180.0)

// Create a new camera
func NewCamera() *Camera {
	return &Camera{
		Camera3D: rl.Camera3D{
			// Position: rl.NewVector3(0, 10, 0),
			// Target:   rl.NewVector3(0, 0, 10),

			Position: rl.NewVector3(0, float32(chunkHeight/2)+10, 0),
			Target:   rl.NewVector3(0, float32(chunkHeight/2), 10),

			Up:         rl.NewVector3(0, 1, 0),
			Fovy:       45.0,
			Projection: rl.CameraPerspective,
		},
	}
}

// TODO: implement
func (c Camera) Frustum() Frustum {

	const renderDistance = 50
	const nearDistance = 0.1

	// forward unit vector (direction)
	forward := rl.Vector3Normalize(
		rl.Vector3Subtract(c.Camera3D.Target, c.Camera3D.Position),
	)

	// Points on the planes
	nearPt := rl.Vector3Add(c.Camera3D.Position,
		rl.Vector3Scale(forward, nearDistance))
	farPt := rl.Vector3Add(c.Camera3D.Position,
		rl.Vector3Scale(forward, renderDistance))

	halfFovyRadians := c.Camera3D.Fovy * (math.Pi / 180.0) / 2.0

	right := rl.Vector3Normalize(rl.Vector3CrossProduct(forward, c.Camera3D.Up))

	topNormal := rl.Vector3Transform(forward,
		// rotate up half fovy then 90 degrees down to get the normal vector
		rl.MatrixRotate(right, halfFovyRadians-ninetyDegRadians),
	)

	bottomNormal := rl.Vector3Transform(forward,
		// rotate down half fovy then 90 degrees up to get the normal vector
		rl.MatrixRotate(right, -halfFovyRadians+ninetyDegRadians),
	)

	rightNormal := rl.Vector3Transform(forward,
		// rotate right half fovy then 90 degrees left to get the normal vector
		rl.MatrixRotate(c.Camera3D.Up, halfFovyRadians-ninetyDegRadians),
	)

	leftNormal := rl.Vector3Transform(forward,
		// rotate left half fovy then 90 degrees right to get the normal vector
		rl.MatrixRotate(c.Camera3D.Up, -halfFovyRadians+ninetyDegRadians),
	)

	/*
	 * equation of a plane
	 * distance: n·P = d
	 * dotProduct(normal vector, point on plane) = distance
	 */

	return Frustum{
		near: Plane{
			normal:   forward,
			distance: rl.Vector3DotProduct(forward, nearPt),
		},
		far: Plane{
			normal:   rl.Vector3Scale(forward, -1),
			distance: -rl.Vector3DotProduct(forward, farPt),
		},
		left: Plane{
			normal:   leftNormal,
			distance: 0,
		},
		right: Plane{
			normal:   rightNormal,
			distance: 0,
		},
		top: Plane{
			normal:   topNormal,
			distance: 0,
		},
		bottom: Plane{
			normal:   bottomNormal,
			distance: 0,
		},
	}
}
