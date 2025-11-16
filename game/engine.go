package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Engine is the main engine struct
type Engine struct {
	*Camera

	World    *World
	Debugger *Debugger
	Input    *InputHandler
}

// Initialize the engine
func NewEngine() (*Engine, error) {
	rl.InitWindow(1000, 800, "Voxel Engine")
	rl.SetTargetFPS(60)

	engine := &Engine{
		World:  NewWorld(),
		Camera: NewCamera(),
	}

	// Initialize the debugger
	debugger := NewDebugger(engine)
	engine.Debugger = &debugger

	// Initialize the input handler
	inputHandler := NewInputHandler(engine)
	engine.Input = &inputHandler

	return engine, nil
}

// GetHorizontalAngleToForward returns the angle between camera vector and world forward vector
func GetHorizontalAngleToForward(cameraVec rl.Vector3) float32 {
	// Create 2D vectors by ignoring Y component
	camDirection2D := rl.Vector2{X: cameraVec.X, Y: cameraVec.Z}
	forward2D := rl.Vector2{X: worldForward.X, Y: worldForward.Z}

	// Normalize the vectors
	camDirection2D = rl.Vector2Normalize(camDirection2D)
	forward2D = rl.Vector2Normalize(forward2D)

	// Calculate angle in radians and convert to degrees
	angle := rl.Rad2deg * rl.Vector2Angle(camDirection2D, forward2D)
	return angle
}

// Main render loop
func (e *Engine) Run() {
	for !rl.WindowShouldClose() {
		e.Input.Handle()
		e.Camera.UpdateFrustum()
		e.render()
	}
	rl.CloseWindow()
}

// Render the world
func (e *Engine) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.BeginMode3D(e.Camera3D)

	// Render chunks
	for _, chunk := range e.World.Chunks {
		chunk.render()
	}

	rl.EndMode3D()

	e.Debugger.Render()

	rl.EndDrawing()
}
