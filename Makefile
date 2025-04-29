# Tidy up dependencies
tidy:
	go mod tidy

# Run the server
run:
	go run cmd/server/main.go

# Build the binary
build:
	go build -o bin/server cmd/server/main.go

# Clean up binaries
clean:
	rm -rf bin/
