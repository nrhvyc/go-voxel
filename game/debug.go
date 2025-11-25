package game

import (
	"fmt"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Debugger renders text about different systems
type Debugger struct {
	engine *Engine

	// Will add config here for what to render
}

func NewDebugger(e *Engine) Debugger {
	return Debugger{
		engine: e,
	}
}

type debugRenderInfo struct {
	chunksRendered []ChunkID
}

func (d Debugger) Render(info debugRenderInfo) {
	rl.DrawFPS(10, 10)

	d.FrustumDebug() // camera frustum
	d.CameraDebug()
	d.ChunkDebug(info.chunksRendered)
}

func (d Debugger) CameraDebug() {
	rl.DrawText(
		fmt.Sprintf(
			"Camera Pos: (%.2f, %.2f, %.2f)",
			d.engine.Camera3D.Position.X,
			d.engine.Camera3D.Position.Y,
			d.engine.Camera3D.Position.Z,
		),
		10, 30, 20, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Target Pos: (%.2f, %.2f, %.2f)",
			d.engine.Camera3D.Target.X,
			d.engine.Camera3D.Target.Y,
			d.engine.Camera3D.Target.Z,
		),
		10, 50, 20, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Chunk Pos: (%#v)",
			d.engine.World.Chunks["0,0,0"].worldPosition,
		),
		10, 70, 20, rl.Black,
	)
}

func (d Debugger) FrustumDebug() {
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.left normal: (%.2f, %.2f, %.2f), distance: (%.2f)",
			d.engine.Camera.Frustum.left.normal.X,
			d.engine.Camera.Frustum.left.normal.Y,
			d.engine.Camera.Frustum.left.normal.Z,
			d.engine.Camera.Frustum.left.distance,
		),
		10, 90, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.right normal: (%.2f, %.2f, %.2f), distance: (%.2f)",
			d.engine.Camera.Frustum.right.normal.X,
			d.engine.Camera.Frustum.right.normal.Y,
			d.engine.Camera.Frustum.right.normal.Z,
			d.engine.Camera.Frustum.right.distance,
		),
		10, 100, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.near normal: (%.2f, %.2f, %.2f), distance: (%.2f)",
			d.engine.Camera.Frustum.near.normal.X,
			d.engine.Camera.Frustum.near.normal.Y,
			d.engine.Camera.Frustum.near.normal.Z,
			d.engine.Camera.Frustum.near.distance,
		),
		10, 110, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.far: (%.2f, %.2f, %.2f), distance: (%.2f)",
			d.engine.Camera.Frustum.far.normal.X,
			d.engine.Camera.Frustum.far.normal.Y,
			d.engine.Camera.Frustum.far.normal.Z,
			d.engine.Camera.Frustum.far.distance,
		),
		10, 120, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.top: (%.2f, %.2f, %.2f), distance: (%.2f)",
			d.engine.Camera.Frustum.top.normal.X,
			d.engine.Camera.Frustum.top.normal.Y,
			d.engine.Camera.Frustum.top.normal.Z,
			d.engine.Camera.Frustum.top.distance,
		),
		10, 130, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.bottom: (%.2f, %.2f, %.2f), distance: (%.2f)",
			d.engine.Camera.Frustum.bottom.normal.X,
			d.engine.Camera.Frustum.bottom.normal.Y,
			d.engine.Camera.Frustum.bottom.normal.Z,
			d.engine.Camera.Frustum.bottom.distance,
		),
		10, 140, 10, rl.Black,
	)
}

func (d Debugger) ChunkDebug(ids []ChunkID) {
	idsStr := make([]string, len(ids))
	for i, id := range ids {
		idsStr[i] = string(id)
	}
	sort.Strings(idsStr)

	rl.DrawText(
		fmt.Sprintf(
			"Chunks Rendered: (%v)",
			idsStr,
		),
		10, 160, 20, rl.Black,
	)

	var chunkTxtPos int32 = 160
	for _, id := range ids {
		chunkTxtPos += 20
		rl.DrawText(
			fmt.Sprintf(
				"Chunk Pos: (%#v)",
				d.engine.World.Chunks[id].worldPosition,
			),
			10, chunkTxtPos, 20, rl.Black,
		)
	}

}
