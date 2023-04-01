compose-dev:
	@echo "[compose-dev]: Running docker compose dev mode..."
	@docker-compose -f infra/docker-compose.yml up --build

compose-stop:
	@echo "[compose-stop]: Stopping docker compose dev mode..."
	@docker-compose -f infra/docker-compose.yml down

backend-dev:
	@echo "[dev] Running Backend..."
	@export $$(cat backend/.env) && cd backend && go run cmd/smartcontracts/main.go

frontend-dev:
	@echo "[dev] Running Frontend..."
	@cd frontend && npm run dev