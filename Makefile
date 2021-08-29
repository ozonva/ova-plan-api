build:
	go build -o ./bin/ ./cmd/ova-plan-api/

run:
	go run ./cmd/ova-plan-api/

test:
	go test ./... -v

generate:
	go generate -v -x ./internal/mockgen.go

proto:
	 protoc -I api/ --go_out=pkg/ova-plan-api --go_opt=paths=import --go-grpc_out=pkg/ova-plan-api --go-grpc_opt=paths=import  ./api/ova-plan-api/api.proto
