run:
	go run main.go
lint:
	golangci-lint run ./...
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o appctl .
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloworld.proto
