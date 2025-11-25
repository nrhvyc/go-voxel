package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Frustum struct {
	left, right, top, bottom, near, far Plane
}

type Plane struct {
	normal   rl.Vector3
	distance float32 // dotProduct(normal vector, point on plane) = distance
}

type Intersection int

const (
	_ Intersection = iota
	Outside
	Inside
	Intersecting
)

// inspired by Real-Time Rendering section 22.10.1
func (p Plane) AABBIntersection(bb rl.BoundingBox) Intersection {
	// get the bbCenter point of the bounding box
	// TODO: calculate these on chunk init
	bbCenter := rl.Vector3Add(bb.Max, bb.Min)
	bbCenter.X, bbCenter.Y, bbCenter.Z = bbCenter.X/2, bbCenter.Y/2, bbCenter.Z/2

	// half diagonal vector from bounding box center point to
	// max point on bounding box
	h := rl.Vector3Subtract(bb.Max, bb.Min)
	h.X, h.Y, h.Z =
		h.X/2, h.Y/2, h.Z/2

	f32Abs := func(n float32) float32 {
		return float32(math.Abs(float64(n)))
	}

	bbExtent := h.X*f32Abs(p.normal.X) + h.Y*f32Abs(p.normal.Y) + h.Z*f32Abs(p.normal.Z)

	signedDistanceToPlane := rl.Vector3DotProduct(bbCenter, p.normal) - p.distance

	if signedDistanceToPlane-bbExtent > 0 {
		return Inside
	} else if signedDistanceToPlane+bbExtent < 0 {
		return Outside
	} else {
		return Intersecting
	}
}

/*
 * equation of a plane
 * distance: n·P = d
 * dotProduct(normal vector, point on plane) = distance
 */

func (f Frustum) Viewable(bb rl.BoundingBox) bool {
	if i := f.left.AABBIntersection(bb); i == Outside {
		return false
	}

	if i := f.right.AABBIntersection(bb); i == Outside {
		return false
	}

	if i := f.top.AABBIntersection(bb); i == Outside {
		return false
	}

	if i := f.bottom.AABBIntersection(bb); i == Outside {
		return false
	}

	// Ignoring the far and near planes for now
	// if i := f.far.AABBIntersection(bb); i == Outside {
	// 	return false
	// }
	// if i := f.near.AABBIntersection(bb); i == Outside {
	// 	return false
	// }

	return true
}
