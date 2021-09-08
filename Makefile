OUTPUT = main 
SERVICE_NAME = drones

.PHONY: test
test:
	go test ./...

build-local:
	go build -o $(OUTPUT) ./cmd/$(SERVICE_NAME)/main.go

run: build-local
	@echo ">> Running application ..."
	PORT=9064 \
	./$(OUTPUT)