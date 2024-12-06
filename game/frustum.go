package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Frustum struct {
	left, right, top, bottom, near, far Plane
}

type Plane struct {
	normal   rl.Vector3
	distance float32 // dotProduct(normal vector, point on plane) = distance
}
