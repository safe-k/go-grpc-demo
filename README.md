# Go gRPC Demo

## Installation

```shell script
git clone git@github.com:seifkamal/go-grpc-demo.git
cd go-grpc-demo
go get -u
```

After cloning the repo, download a suitable prebuilt binary version of the compiler from 
https://github.com/protocolbuffers/protobuf/releases and compile the `proto/service.proto` definition file.

**Example:**
```shell script
# --proto_path=proto - Directory where the proto definition is located
# --proto_path=include - Directory that comes with the compiler (used to generate the Go code)
# --go_out=plugins=grpc:proto - Use protoc-gen-go to generate code compatible with gRPC (see https://github.com/golang/protobuf#grpc-support)
protoc --proto_path=proto --proto_path=include --go_out=plugins=grpc:proto service.proto
```

## Usage

### Server

Run a server instance:
```shell script
go run server.go
```

This will open a TCP connection over port 4040.

### Client

You can connect to the server via a client of your choice, or optionally run a client instance:
```shell script
go run client/client.go
```

This will start an HTTP server on port 8080, with the following endpoints:
- `/add/:a/:b`
- `/multiply/:a/:b`
