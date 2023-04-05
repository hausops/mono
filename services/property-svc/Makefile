codegen:
	protoc --go_out=. --go-grpc_out=. pb/*.proto

dev:
	go run cmd/server/server.go -c config/config.dev.yaml
