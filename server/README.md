## Build

```
go build ./...
```

## Run

Stateless variant:

```
go run main.go --port 8889 stateless
```

Stateful variant:

```
go run main.go --port 8889 stateful
```

## gRPC

Generate gRPC client/server code:

```
protoc --proto_path ../protos --go_out=plugins=grpc:stateless/numbers ../protos/numbers.proto
```

```
protoc --proto_path ../protos --go_out=plugins=grpc:stateful/randstream ../protos/randstream.proto
```