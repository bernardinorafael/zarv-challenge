.PHONY: run
run: # Execute the program
	@go run cmd/main.go

.PHONY: test
test: # Run the tests
	@go test -v ./...
