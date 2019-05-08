## Build

```
go build ./...
```

## Run

```
go run main.go -port 8888
```

## gRPC

Generate gRPC client/server code:

```
protoc --proto_path ../protos --go_out=plugins=grpc:numbers ../protos/numbers.proto
```
