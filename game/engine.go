package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	mouseRotateSpeed = 0.003
)

// Camera represents the player's view
type Camera struct {
	Camera3D rl.Camera3D
}

// Engine is the main engine struct
type Engine struct {
	*Camera

	World *World
}

// Vertex data for a single cube
var cubeVertices = []float32{
	// Front face
	-0.5, -0.5, 0.5,
	0.5, -0.5, 0.5,
	0.5, 0.5, 0.5,
	0.5, 0.5, 0.5,
	-0.5, 0.5, 0.5,
	-0.5, -0.5, 0.5,

	// Back face
	-0.5, -0.5, -0.5,
	-0.5, 0.5, -0.5,
	0.5, 0.5, -0.5,
	0.5, 0.5, -0.5,
	0.5, -0.5, -0.5,
	-0.5, -0.5, -0.5,

	// Top face
	-0.5, 0.5, -0.5,
	-0.5, 0.5, 0.5,
	0.5, 0.5, 0.5,
	0.5, 0.5, 0.5,
	0.5, 0.5, -0.5,
	-0.5, 0.5, -0.5,

	// Bottom face
	-0.5, -0.5, -0.5,
	0.5, -0.5, -0.5,
	0.5, -0.5, 0.5,
	0.5, -0.5, 0.5,
	-0.5, -0.5, 0.5,
	-0.5, -0.5, -0.5,

	// Right face
	0.5, -0.5, -0.5,
	0.5, 0.5, -0.5,
	0.5, 0.5, 0.5,
	0.5, 0.5, 0.5,
	0.5, -0.5, 0.5,
	0.5, -0.5, -0.5,

	// Left face
	-0.5, -0.5, -0.5,
	-0.5, -0.5, 0.5,
	-0.5, 0.5, 0.5,
	-0.5, 0.5, 0.5,
	-0.5, 0.5, -0.5,
	-0.5, -0.5, -0.5,
}

// Initialize the engine
func NewEngine() (*Engine, error) {
	rl.InitWindow(800, 600, "Voxel Engine")
	rl.SetTargetFPS(60)

	engine := &Engine{
		World:  NewWorld(),
		Camera: NewCamera(),
	}

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
		e.handleInput()
		e.render()
	}
	rl.CloseWindow()
}

// Handle keyboard input
func (e *Engine) handleInput() {
	speed := float32(0.005)
	if rl.IsKeyDown(rl.KeyLeftShift) {
		speed *= 2
	}

	// Forward
	if rl.IsKeyDown(rl.KeyW) {
		e.Camera3D.Position = rl.Vector3Add(
			e.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
				speed,
			),
		)
		e.Camera3D.Target = rl.Vector3Add(
			e.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
				speed,
			),
		)
		// This moves the camera forward by:
		// 1. Calculating the direction vector (Target - Position)
		// 2. Scaling this vector by the speed
		// 3. Adding the scaled vector to the current position
	}

	// Backward
	if rl.IsKeyDown(rl.KeyS) {
		e.Camera3D.Position = rl.Vector3Subtract(
			e.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
				speed,
			),
		)
		e.Camera3D.Target = rl.Vector3Subtract(
			e.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
				speed,
			),
		)
		// This moves the camera backward by:
		// 1. Calculating the direction vector (Target - Position)
		// 2. Scaling this vector by the speed
		// 3. Subtracting the scaled vector from the current position
	}

	// Left
	if rl.IsKeyDown(rl.KeyA) {
		e.Camera3D.Position = rl.Vector3Subtract(
			e.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
					e.Camera3D.Up,
				),
				speed,
			),
		)
		e.Camera3D.Target = rl.Vector3Subtract(
			e.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
					e.Camera3D.Up,
				),
				speed,
			),
		)
	}

	// Right
	if rl.IsKeyDown(rl.KeyD) {
		e.Camera3D.Position = rl.Vector3Add(
			e.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
					e.Camera3D.Up,
				),
				speed,
			),
		)
		e.Camera3D.Target = rl.Vector3Add(
			e.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
					e.Camera3D.Up,
				),
				speed,
			),
		)
	}

	// Up
	if rl.IsKeyDown(rl.KeySpace) {
		upVector := rl.Vector3Scale(e.Camera3D.Up, speed*10)
		e.Camera3D.Position = rl.Vector3Add(e.Camera3D.Position, upVector)
		e.Camera3D.Target = rl.Vector3Add(e.Camera3D.Target, upVector)
	}

	// Down
	if rl.IsKeyDown(rl.KeyLeftAlt) {
		downVector := rl.Vector3Scale(e.Camera3D.Up, -speed*10)
		e.Camera3D.Position = rl.Vector3Add(e.Camera3D.Position, downVector)
		e.Camera3D.Target = rl.Vector3Add(e.Camera3D.Target, downVector)
	}

	// Mouse Camera rotation
	mousePositionDelta := rl.GetMouseDelta()
	mouseInvertOption := false // TODO: extract as config option
	var mouseInvert float32 = 1
	if mouseInvertOption {
		mouseInvert = -1
	}

	// Horizontal camera rotation
	e.Camera3D.Target = rl.Vector3Add(
		e.Camera3D.Position,
		rl.Vector3Transform(
			rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position),
			rl.MatrixRotateY(mouseInvert*mousePositionDelta.X*mouseRotateSpeed),
		),
	)

	// Vertical rotation
	cameraVec := rl.Vector3Subtract(e.Camera3D.Target, e.Camera3D.Position)
	right := rl.Vector3CrossProduct(cameraVec, e.Camera3D.Up)
	right = rl.Vector3Normalize(right)

	// Create rotation matrix around right vector
	rotationMatrix := rl.MatrixRotate(right, -1*mouseInvert*mousePositionDelta.Y*mouseRotateSpeed)

	newTarget := rl.Vector3Add(
		e.Camera3D.Position,
		rl.Vector3Transform(cameraVec, rotationMatrix),
	)

	// Calculate vertical angle for clamping
	camDirection := rl.Vector3Subtract(newTarget, e.Camera3D.Position)
	angleVertical := rl.Rad2deg * rl.Vector3Angle(camDirection, e.Camera3D.Up)

	// Clamp vertical rotation between 20 and 160 degrees
	if angleVertical >= 20 && angleVertical <= 160 {
		e.Camera3D.Target = newTarget
	}

	// Hide cursor and center it
	rl.HideCursor()
	rl.SetMousePosition(rl.GetScreenWidth()/2, rl.GetScreenHeight()/2)
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

	// Draw debug text
	rl.DrawFPS(10, 10)
	rl.DrawText(
		fmt.Sprintf(
			"Camera Pos: (%.2f, %.2f, %.2f)",
			e.Camera3D.Position.X,
			e.Camera3D.Position.Y,
			e.Camera3D.Position.Z,
		),
		10, 30, 20, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Target Pos: (%.2f, %.2f, %.2f)",
			e.Camera3D.Target.X,
			e.Camera3D.Target.Y,
			e.Camera3D.Target.Z,
		),
		10, 50, 20, rl.Black,
	)
	rl.DrawText(
		fmt.Sprintf(
			"Chunk Pos: (%#v)",
			e.World.Chunks["0,0,0"].Position,
		),
		10, 70, 20, rl.Black,
	)

	rl.EndDrawing()
}
