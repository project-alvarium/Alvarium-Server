.PHONY: build clean run test

build:
	CGO_ENABLED=1 go build -o database-manager

clean:
	rm ./database-manager

run:
	cd bin && ./launch.sh

test:
	go test -coverprofile=coverage.out ./...
	go vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]