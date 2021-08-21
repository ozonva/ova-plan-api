build:
	go build -o ./bin/ ./cmd/ova-plan-api/

run:
	go run ./cmd/ova-plan-api/

test:
	go test ./... -v

generate:
	go generate ./...