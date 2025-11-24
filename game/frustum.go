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

	signedDistanceToPlane := rl.Vector3DotProduct(bbCenter, p.normal)

	if signedDistanceToPlane > bbExtent {
		return Outside
	} else if signedDistanceToPlane < bbExtent {
		return Inside
	} else {
		return Intersecting
	}
}

/*
 * equation of a plane
 * distance: n·P = d
 * dotProduct(normal vector, point on plane) = distance
 */

// TODO: implement
func (f Frustum) Intersection(bb rl.BoundingBox) bool {
	// check side planes first

	// boundingBoxLeftPlan rl.Vector3{X:1, Y:0, Z:0}

	return true
}
