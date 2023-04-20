.PHONY: help

.DEFAULT_GOAL := help

go-fumpt: ## TODO: weigh out between go-fmt, goimports, and go-fumpt. also golangci-lint and govet (govet in lint?)
	@gofumpt

compose-up: ## spins up the stack in compose.yaml; builds images (if needed), and runs the containers in the background
	@docker compose up -d --build --remove-orphans

compose-logs: ## attaches to the logs output of the repo's compose stack (can be quit with (ctrl/cmd)-C)
	@docker compose logs -t -f --tail 250

compose-stop: ## stops all running containers in the repo's compose stack
	@docker compose stop

compose-down: ## spin docker compose stack down, removing all containers and networks
	@docker compose down --remove-orphans

migrate-new: ## creates a new sql goose migration file, e.g., make goose-new name=add_some_column_to_some_table
	@docker compose run tool goose -dir migrations create $(name) sql

migrate-fix: ## fixes timestamped sql migration names to a sequential version number (according to dir contents)
	@docker compose run tool goose -dir migrations fix

migrate-up: ## runs goose migrations up (by version number)
	@goose -dir migrations postgres $(GOOSE_DBSTRING) up

migrate-down: ## runs goose migrations down (by version number)
	@goose -dir migrations postgres $(GOOSE_DBSTRING) down

help: ## prints out the help documentation (also will be printed by simply running `make` command with no arg)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
