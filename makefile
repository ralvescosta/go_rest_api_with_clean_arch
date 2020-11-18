run-dev:
	go run ./src/*.go

test:
	go test ./src/...

test-cov:
	go test ./src/... -cover -v -coverprofile=c.out && go tool cover -html=c.out -o coverage.html

build:
	go build ./src/*.go