migrate:
	go run src/frameworks/database/migration/main.go

run:
	go run src/main.go

tests:
	go test ./... -v

test-cov:
	go test ./... -cover -v -coverprofile=c.out && go tool cover -html=c.out -o coverage.html

build:
	go build src/main.go