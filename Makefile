.PHONY: build
build:
	make beauty && mkdir -p bin; go build -o bin/database-mesh pkg/database-mesh.go
beauty:
	gofmt -w .
run:
	make build && LOGLEVEL=INFO bin/database-mesh
clean:
	rm -r bin
debug:
	make build && LOGLEVEL=DEBUG bin/database-mesh
linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/database-mesh pkg/database-mesh.go
	echo `git log | grep commit | head -1 | cut -d" " -f2` > manifest
docker:
	make linux && docker build -t mworks92/database-mesh:`date +%Y%m%d%H%M%S` .
lint:
	golint ./...
