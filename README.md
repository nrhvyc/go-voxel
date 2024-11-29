# go-voxel
A voxel engine written using Go &amp; Vulkan


## Work in Progress Compiling to WASM
GOOS=js GOARCH=wasm go build -o main.wasm cmd/*

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux" CXX="zig c++ -target x86_64-linux" go build -o main.wasm cmd/*

CGO_ENABLED=1 GOOS=js GOARCH=wasm CC="zig cc -target x86_64-linux" CXX="zig c++ -target x86_64-linux" go build -o main.wasm cmd/*

