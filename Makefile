fmt:
	docker compose exec app gofumpt -l -w .
lint:
	docker compose exec app golangci-lint run
