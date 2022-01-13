doc:
	@echo "Creating docs templates"
	@swag init -g cmd/api/main.go -o cmd/api/docs

linter:
	@echo "Linter running"
	@golangci-lint run ./...
