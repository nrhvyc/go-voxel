package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ninetyDegRadians = 90.0 * (math.Pi / 180.0)

const frustumRenderDistance = 50
const frustumNearDistance = 0.1

// Camera represents the player's view
type Camera struct {
	Camera3D rl.Camera3D

	Frustum Frustum // currently in world space
}

// Create a new camera
func NewCamera() *Camera {
	camera := Camera{
		Camera3D: rl.Camera3D{
			// Position: rl.NewVector3(0, 10, 0),
			// Target:   rl.NewVector3(0, 0, 10),

			Position: rl.NewVector3(0, float32(chunkHeight/2)+10, 0),
			Target:   rl.NewVector3(0, float32(chunkHeight/2), 10),

			Up:         rl.NewVector3(0, 1, 0),
			Fovy:       45.0,
			Projection: rl.CameraPerspective,
		},
		Frustum: Frustum{},
	}

	camera.UpdateFrustum()

	return &camera
}

func (c *Camera) UpdateFrustum() {

	// forward unit vector (direction)
	forward := rl.Vector3Normalize(
		rl.Vector3Subtract(c.Camera3D.Target, c.Camera3D.Position),
	)

	// Points on the planes
	nearPt := rl.Vector3Add(c.Camera3D.Position,
		rl.Vector3Scale(forward, frustumNearDistance))
	farPt := rl.Vector3Add(c.Camera3D.Position,
		rl.Vector3Scale(forward, frustumRenderDistance))

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

	// FOV differs for width depending on the screen width
	aspectRatio := float32(rl.GetScreenWidth()) / float32(rl.GetScreenHeight())
	halfFovxRadians := math.Atan(math.Tan(float64(halfFovyRadians)) * float64(aspectRatio))

	rightNormal := rl.Vector3Transform(forward,
		// rotate right half fovy then 90 degrees left to get the normal vector
		rl.MatrixRotate(c.Camera3D.Up, float32(halfFovxRadians)-ninetyDegRadians),
	)

	leftNormal := rl.Vector3Transform(forward,
		// rotate left half fovy then 90 degrees right to get the normal vector
		rl.MatrixRotate(c.Camera3D.Up, -float32(halfFovxRadians)+ninetyDegRadians),
	)

	// TODO: convert the frustum to view space later since apparently
	// that's more efficient to check, but world space is easier
	// for me to reason about for now

	/*
	 * equation of a plane
	 * distance: n·P = d
	 * dotProduct(normal vector, point on plane) = distance
	 */

	// world space
	c.Frustum = Frustum{
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
			distance: rl.Vector3DotProduct(leftNormal, c.Camera3D.Position),
		},
		right: Plane{
			normal:   rightNormal,
			distance: rl.Vector3DotProduct(rightNormal, c.Camera3D.Position),
		},
		top: Plane{
			normal:   topNormal,
			distance: rl.Vector3DotProduct(topNormal, c.Camera3D.Position),
		},
		bottom: Plane{
			normal:   bottomNormal,
			distance: rl.Vector3DotProduct(bottomNormal, c.Camera3D.Position),
		},
	}
}
