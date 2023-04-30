build:
	go build -o bin/dashboard-api cmd/graphql/main.go

codegen:
	go generate graphql/resolver.go

run:
	dotenv go run cmd/graphql/main.go
