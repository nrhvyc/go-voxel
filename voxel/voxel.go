package voxel

import (
	"fmt"
	"strings"
	// "log/slog"
	// "os"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// var handler slog.Handler = slog.NewTextHandler(os.Stdout, nil)
// var logger = slog.New(handler)

// Voxel represents a single cube in the world
type Voxel struct {
	Position mgl32.Vec3
	Type     int
}

// Chunk represents a 16x16x16 section of voxels
type Chunk struct {
	Voxels   [16][16][16]*Voxel
	Position mgl32.Vec3
}

// World contains all chunks
type World struct {
	Chunks map[string]*Chunk
}

// Camera represents the player's view
type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
}

// Engine is the main engine struct
type Engine struct {
	Window        *glfw.Window
	World         *World
	Camera        *Camera
	ShaderProgram uint32
	VAO           uint32 // vertex array object
	VBO           uint32 // vertex buffer object
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
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		return nil, fmt.Errorf("Failed to initialize GLFW:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(800, 600, "Voxel Engine", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create window:", err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, fmt.Errorf("Failed to initialize OpenGL:", err)
	}

	engine := &Engine{
		Window: window,
		World:  NewWorld(),
		Camera: NewCamera(),
	}

	engine.initShaders()
	engine.initBuffers()

	// Enable depth testing
	gl.Enable(gl.DEPTH_TEST)

	return engine, nil
}

// Initialize OpenGL buffers for the cube mesh
func (e *Engine) initBuffers() {
	// Generate and bind VAO
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	e.VAO = vao

	// Generate and bind VBO
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	e.VBO = vbo

	// Upload vertex data
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

	// Set vertex attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	// Unbind VAO and VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

// Create a new camera
func NewCamera() *Camera {
	return &Camera{
		Position: mgl32.Vec3{0, 0, 20},
		Front:    mgl32.Vec3{0, 0, -1},
		Up:       mgl32.Vec3{0, 1, 0},
	}
}

// Create a new world
func NewWorld() *World {
	world := &World{
		Chunks: make(map[string]*Chunk),
	}

	// Add a single chunk with some voxels
	chunk := &Chunk{
		Position: mgl32.Vec3{0, 0, 0},
	}
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				if x == 0 || y == 0 || z == 0 || x == 15 || y == 15 || z == 15 {
					chunk.Voxels[x][y][z] = &Voxel{
						Position: mgl32.Vec3{float32(x), float32(y), float32(z)},
						Type:     1,
					}
				}
			}
		}
	}
	world.Chunks["0,0,0"] = chunk

	return world
}

// Initialize shaders
func (e *Engine) initShaders() {
	// GLSL vertex shader
	vertexShader := `
		#version 410
		layout (location = 0) in vec3 position;
		layout (location = 1) in vec3 color;
		out vec3 fragColor;
		uniform mat4 projection;
		uniform mat4 view;
		uniform mat4 model;
		void main() {
			gl_Position = projection * view * model * vec4(position, 1.0);
			fragColor = color;
		}
	`
	// GLSL fragment shader
	fragmentShader := `
		#version 410
		in vec3 fragColor;
		out vec4 FragColor;
		void main() {
			FragColor = vec4(fragColor, 1.0);
		}
    `

	e.ShaderProgram = createShaderProgram(vertexShader, fragmentShader)
}

// Main render loop
func (e *Engine) Run() {
	for !e.Window.ShouldClose() {
		e.handleInput()
		e.render()
		e.Window.SwapBuffers()
		glfw.PollEvents()
	}
}

// Handle keyboard input
func (e *Engine) handleInput() {
	if e.Window.GetKey(glfw.KeyEscape) == glfw.Press {
		e.Window.SetShouldClose(true)
	}

	speed := float32(0.1)
	if e.Window.GetKey(glfw.KeyLeftShift) == glfw.Press {
		speed *= 2
	}

	// Forward/Backward
	if e.Window.GetKey(glfw.KeyW) == glfw.Press {
		e.Camera.Position = e.Camera.Position.Add(e.Camera.Front.Mul(speed))
	}
	if e.Window.GetKey(glfw.KeyS) == glfw.Press {
		e.Camera.Position = e.Camera.Position.Sub(e.Camera.Front.Mul(speed))
	}

	// Left/Right
	if e.Window.GetKey(glfw.KeyA) == glfw.Press {
		e.Camera.Position = e.Camera.Position.Sub(e.Camera.Front.Cross(e.Camera.Up).Normalize().Mul(speed))
	}
	if e.Window.GetKey(glfw.KeyD) == glfw.Press {
		e.Camera.Position = e.Camera.Position.Add(e.Camera.Front.Cross(e.Camera.Up).Normalize().Mul(speed))
	}

	// Up/Down
	if e.Window.GetKey(glfw.KeySpace) == glfw.Press {
		e.Camera.Position = e.Camera.Position.Add(e.Camera.Up.Mul(speed))
	}
	if e.Window.GetKey(glfw.KeyLeftShift) == glfw.Press {
		e.Camera.Position = e.Camera.Position.Sub(e.Camera.Up.Mul(speed))
	}
}

// Render the world
func (e *Engine) render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(e.ShaderProgram)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), 800.0/600.0, 0.1, 100.0)
	view := mgl32.LookAtV(e.Camera.Position, e.Camera.Position.Add(e.Camera.Front), e.Camera.Up)

	projectionLoc := gl.GetUniformLocation(e.ShaderProgram, gl.Str("projection\x00"))
	viewLoc := gl.GetUniformLocation(e.ShaderProgram, gl.Str("view\x00"))

	gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])
	gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])

	// Bind VAO before rendering
	gl.BindVertexArray(e.VAO)

	// Render chunks
	for _, chunk := range e.World.Chunks {
		e.renderChunk(chunk)
	}

	// Unbind VAO after rendering
	gl.BindVertexArray(0)
}

// Render a single chunk
func (e *Engine) renderChunk(chunk *Chunk) {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				if voxel := chunk.Voxels[x][y][z]; voxel != nil {
					model := mgl32.Translate3D(
						chunk.Position.X()+float32(x),
						chunk.Position.Y()+float32(y),
						chunk.Position.Z()+float32(z),
					).Mul4(mgl32.Scale3D(0.9, 0.9, 0.9))
					modelLoc := gl.GetUniformLocation(e.ShaderProgram, gl.Str("model\x00"))
					gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

					// Set color based on voxel type or position
					r := float32(x) / 16.0
					g := float32(y) / 16.0
					b := float32(z) / 16.0
					colorLoc := gl.GetUniformLocation(e.ShaderProgram, gl.Str("color\x00"))
					gl.Uniform3f(colorLoc, r, g, b)

					gl.DrawArrays(gl.TRIANGLES, 0, 36)
				}
			}
		}
	}
}

// Helper function to create shader program
func createShaderProgram(vertexSrc, fragmentSrc string) uint32 {
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)

	csources, free := gl.Strs(vertexSrc + "\x00")
	gl.ShaderSource(vertexShader, 1, csources, nil)
	free()

	csources, free = gl.Strs(fragmentSrc + "\x00")
	gl.ShaderSource(fragmentShader, 1, csources, nil)
	free()

	gl.CompileShader(vertexShader)
	if err := getShaderCompileError(vertexShader); err != nil {
		fmt.Printf("Vertex shader compilation failed: %v\n", err)
		return 0
	}
	fmt.Println("Vertex shader compiled successfully")

	gl.CompileShader(fragmentShader)
	if err := getShaderCompileError(fragmentShader); err != nil {
		fmt.Printf("Fragment shader compilation failed: %v\n", err)
		return 0
	}
	fmt.Println("Fragment shader compiled successfully")

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var success int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
		fmt.Printf("Program linking failed: %v\n", log)
		return 0
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

func getShaderCompileError(shader uint32) error {
	var success int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return fmt.Errorf(log)
	}
	return nil
}

// Clean up resources
func (e *Engine) Cleanup() {
	gl.DeleteVertexArrays(1, &e.VAO)
	gl.DeleteBuffers(1, &e.VBO)
	gl.DeleteProgram(e.ShaderProgram)
}
