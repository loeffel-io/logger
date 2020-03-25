install:
	go get

linter:
	golangci-lint run --timeout=2m