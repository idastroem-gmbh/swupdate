#!/bin/bash
docker build -t protoc-go:latest .
docker run --rm -v $(pwd):/workspace -w /workspace protoc-go:latest \
  --proto_path=/workspace \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative swupdate.proto
