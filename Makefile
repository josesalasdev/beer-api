doc:
	@echo "Creando docs templates"
	@swag init -g cmd/api/main.go -o cmd/api/docs
	@cp cmd/api/docs/swagger.yaml .
