run/api:
	@echo "Running API"
	@go run cmd/api

dev/api:
	@echo "Running API in development mode"
	@air -c .air.toml