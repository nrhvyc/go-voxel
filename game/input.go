package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	mouseRotateSpeed = 0.003
)

type InputHandler struct {
	engine *Engine
}

func NewInputHandler(e *Engine) InputHandler {
	return InputHandler{
		engine: e,
	}
}

// Handle keyboard input
func (ih *InputHandler) Handle() {
	speed := float32(0.005)
	if rl.IsKeyDown(rl.KeyLeftShift) {
		speed *= 2
	}

	// Forward
	if rl.IsKeyDown(rl.KeyW) {
		ih.engine.Camera3D.Position = rl.Vector3Add(
			ih.engine.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
				speed,
			),
		)
		ih.engine.Camera3D.Target = rl.Vector3Add(
			ih.engine.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
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
		ih.engine.Camera3D.Position = rl.Vector3Subtract(
			ih.engine.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
				speed,
			),
		)
		ih.engine.Camera3D.Target = rl.Vector3Subtract(
			ih.engine.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
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
		ih.engine.Camera3D.Position = rl.Vector3Subtract(
			ih.engine.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
					ih.engine.Camera3D.Up,
				),
				speed,
			),
		)
		ih.engine.Camera3D.Target = rl.Vector3Subtract(
			ih.engine.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
					ih.engine.Camera3D.Up,
				),
				speed,
			),
		)
	}

	// Right
	if rl.IsKeyDown(rl.KeyD) {
		ih.engine.Camera3D.Position = rl.Vector3Add(
			ih.engine.Camera3D.Position,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
					ih.engine.Camera3D.Up,
				),
				speed,
			),
		)
		ih.engine.Camera3D.Target = rl.Vector3Add(
			ih.engine.Camera3D.Target,
			rl.Vector3Scale(
				rl.Vector3CrossProduct(
					rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
					ih.engine.Camera3D.Up,
				),
				speed,
			),
		)
	}

	// Up
	if rl.IsKeyDown(rl.KeySpace) {
		upVector := rl.Vector3Scale(ih.engine.Camera3D.Up, speed*10)
		ih.engine.Camera3D.Position = rl.Vector3Add(ih.engine.Camera3D.Position, upVector)
		ih.engine.Camera3D.Target = rl.Vector3Add(ih.engine.Camera3D.Target, upVector)
	}

	// Down
	if rl.IsKeyDown(rl.KeyLeftAlt) {
		downVector := rl.Vector3Scale(ih.engine.Camera3D.Up, -speed*10)
		ih.engine.Camera3D.Position = rl.Vector3Add(ih.engine.Camera3D.Position, downVector)
		ih.engine.Camera3D.Target = rl.Vector3Add(ih.engine.Camera3D.Target, downVector)
	}

	// Mouse Camera rotation
	mousePositionDelta := rl.GetMouseDelta()
	mouseInvertOption := false // TODO: extract as config option
	var mouseInvert float32 = 1
	if mouseInvertOption {
		mouseInvert = -1
	}

	// Horizontal camera rotation
	ih.engine.Camera3D.Target = rl.Vector3Add(
		ih.engine.Camera3D.Position,
		rl.Vector3Transform(
			rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position),
			rl.MatrixRotateY(mouseInvert*mousePositionDelta.X*mouseRotateSpeed),
		),
	)

	// Vertical rotation
	cameraVec := rl.Vector3Subtract(ih.engine.Camera3D.Target, ih.engine.Camera3D.Position)
	right := rl.Vector3CrossProduct(cameraVec, ih.engine.Camera3D.Up)
	right = rl.Vector3Normalize(right)

	// Create rotation matrix around right vector
	rotationMatrix := rl.MatrixRotate(right, -1*mouseInvert*mousePositionDelta.Y*mouseRotateSpeed)

	newTarget := rl.Vector3Add(
		ih.engine.Camera3D.Position,
		rl.Vector3Transform(cameraVec, rotationMatrix),
	)

	// Calculate vertical angle for clamping
	camDirection := rl.Vector3Subtract(newTarget, ih.engine.Camera3D.Position)
	angleVertical := rl.Rad2deg * rl.Vector3Angle(camDirection, ih.engine.Camera3D.Up)

	// Clamp vertical rotation between 20 and 160 degrees
	if angleVertical >= 20 && angleVertical <= 160 {
		ih.engine.Camera3D.Target = newTarget
	}

	// Hide cursor and center it
	rl.HideCursor()
	rl.SetMousePosition(rl.GetScreenWidth()/2, rl.GetScreenHeight()/2)
}
