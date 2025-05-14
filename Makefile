lint:
	golangci-lint run --fix

install:
	go install ./cmd/oas-combiner
