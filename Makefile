build:
	go build -o bin/ova-plan-api ./cmd/ova-plan-api/main.go

run:
	go run ./cmd/ova-plan-api/main.go

test:
	go test ./...

test-integration:
	go test --tags=integration -count 1 ./integration_test/...

generate:
	go generate -v -x ./internal/mockgen.go

proto:
	protoc -I api/ --go_out=pkg/ova-plan-api --go_opt=paths=import --go-grpc_out=pkg/ova-plan-api --go-grpc_opt=paths=import  ./api/ova-plan-api/api.proto

migrations-run:
	goose -dir=migrations postgres "postgresql://${OVA_PLAN_DB_USER}:${OVA_PLAN_DB_PASSWORD}@${OVA_PLAN_DB_HOST}:${OVA_PLAN_DB_PORT}/${OVA_PLAN_DB_NAME}?sslmode=disable" up

run-with-infra:
	docker compose up

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
	go get -u github.com/pressly/goose/v3/cmd/goose
	go get -u github.com/onsi/ginkgo@v1.16.4
	go get -u github.com/onsi/gomega@v1.16.0
	go get -u github.com/golang/mock@v1.6.0
	go get -u github.com/rs/zerolog/log@v1.23.0
	go get -d github.com/pressly/goose/v3/cmd/goose@v3.1.0

prepare: build