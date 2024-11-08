FROM golang:1.23.2

# Install necessary tools
RUN apt-get update && \
    apt-get install -y unzip && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.29.1 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
    export PATH="$PATH:$(go env GOPATH)/bin"

# Download and install the latest protoc
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v28.3/protoc-28.3-linux-x86_64.zip && \
    unzip protoc-28.3-linux-x86_64.zip -d /usr/local && \
    rm protoc-28.3-linux-x86_64.zip

WORKDIR /workspace
ENTRYPOINT ["protoc"]
