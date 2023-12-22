# Link

- [YouTube Link](https://www.youtube.com/watch?v=a6G5-LUlFO4)

## Install protoc

- For Mac, use homebrew to install protoc `brew install protobuf`, `protoc --version`
- [Download Link for gRPC go](https://grpc.io/docs/languages/go/quickstart/)

  ```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

  export PATH="$PATH:$(go env GOPATH)/bin"
  ```
- Command to generate 2 `.pb.go` files
    - Run `protoc --go_out=. --go-grpc_out=. proto/greet.proto`
