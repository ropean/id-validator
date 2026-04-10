# note: call scripts from /scripts
GOPATH=/opt/homebrew/bin/go=go

dev:
	DEV=1 go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

start:
	./bin/server

docker-build:
	docker compose build

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

golangci-lint:
	golangci-lint run ./... --color=always --verbose

gosec:
	gosec ./... -color -verbose

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -cover -coverprofile=coverage -covermode=atomic -race -v

bench:
	go test ./... -bench=. -benchmem -v

cover:
	go tool cover -html=coverage
