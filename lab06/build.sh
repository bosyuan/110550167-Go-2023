cd wasm
go env -w GOOS=js GOARCH=wasm
go build -o lib.wasm main.go 
cd ..
go env -w GOOS=windows GOARCH=amd64
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"