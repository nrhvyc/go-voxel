package voxel

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	mouseRotateSpeed = 0.003
)

var (
	worldUp      = rl.Vector3{X: 0, Y: 1, Z: 0}
	worldForward = rl.Vector3{X: 0, Y: 0, Z: 1}
)

// var handler slog.Handler = slog.NewTextHandler(os.Stdout, nil)
// var logger = slog.New(handler)

// Voxel represents a single cube in the world
type Voxel struct {
	Position rl.Vector3
	Type     int
}

// Chunk represents a 16x16x16 section of voxels
type Chunk struct {
	Voxels   [16][16][16]*Voxel
	Position rl.Vector3
}

// World contains all chunks
type World struct {
	Chunks map[string]*Chunk
}

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

// Create a new camera
func NewCamera() *Camera {
	return &Camera{
		Camera3D: rl.Camera3D{
			Position:   rl.NewVector3(0, 0, 20),
			Target:     rl.NewVector3(0, 0, 0),
			Up:         rl.NewVector3(0, 1, 0),
			Fovy:       45.0,
			Projection: rl.CameraPerspective,
		},
	}
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

// Create a new world
func NewWorld() *World {
	world := &World{
		Chunks: make(map[string]*Chunk),
	}

	// Add a single chunk with some voxels
	chunk := &Chunk{
		Position: rl.NewVector3(0, 0, 0),
	}
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				if x == 0 || y == 0 || z == 0 || x == 15 || y == 15 || z == 15 {
					chunk.Voxels[x][y][z] = &Voxel{
						Position: rl.NewVector3(float32(x), float32(y), float32(z)),
						Type:     1,
					}
				}
			}
		}
	}
	world.Chunks["0,0,0"] = chunk

	return world
}

// No need for initShaders function in Raylib

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
		e.renderChunk(chunk)
	}

	rl.EndMode3D()
	rl.EndDrawing()
}

// Render a single chunk
func (e *Engine) renderChunk(chunk *Chunk) {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				if voxel := chunk.Voxels[x][y][z]; voxel != nil {
					position := rl.Vector3Add(chunk.Position, rl.NewVector3(float32(x), float32(y), float32(z)))
					color := rl.NewColor(uint8(x*16), uint8(y*16), uint8(z*16), 255)
					rl.DrawCube(position, 0.9, 0.9, 0.9, color)
				}
			}
		}
	}
}
