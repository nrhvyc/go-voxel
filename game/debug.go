package game

import (
	"fmt"

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

func (d Debugger) Render() {
	rl.DrawFPS(10, 10)

	d.drawFrustumDebug()
	d.drawCameraDebug()
}

func (d Debugger) drawFrustumDebug() {
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.left: (%v)",
			d.engine.Camera.Frustum.left,
		),
		10, 90, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.right: (%v)",
			d.engine.Camera.Frustum.right,
		),
		10, 100, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.near: (%v)",
			d.engine.Camera.Frustum.near,
		),
		10, 110, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.far: (%v)",
			d.engine.Camera.Frustum.far,
		),
		10, 120, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.top: (%v)",
			d.engine.Camera.Frustum.top,
		),
		10, 130, 10, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Frustum.bottom: (%.v)",
			d.engine.Camera.Frustum.bottom,
		),
		10, 140, 10, rl.Black,
	)
}

func (d Debugger) drawCameraDebug() {
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
