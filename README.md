### grpc

protoc is from protoc-gen-go (installed via go and need to be on path)
Call "protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/_.proto"
or use alias "protoc ./proto/_.proto"

After generating the code run "go mod tidy" to update the go.mod file with the new dependencies

if client gets something like "Unimplemented desc = method ExampleMethod not implemented",
the go server might be missing the implementation of the method
