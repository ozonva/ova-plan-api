LOCAL_BIN:=$(CURDIR)/bin

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

deps: .install-go-deps

.install-go-deps:
	ls go.mod || go mod init github.com/ozonva/ova-plan-api
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger