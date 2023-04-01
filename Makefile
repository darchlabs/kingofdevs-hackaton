backend-dev:
	@echo "[dev] Running Backend..."
	@export $$(cat backend/.env) && cd backend && go run cmd/smartcontracts/main.go

frontend-dev:
	@echo "[dev] Running Frontend..."
	@cd frontend && npm run dev

build-back:
	@echo "[building backend]"
	@docker build -t darchlabs/backend-hackathon -f ./backend/Dockerfile --progress tty .
	@echo "Build darchlabs/backend-hackathon docker image done ✔︎"

build-front:
	@echo "[building frontend]"
	@docker build -t darchlabs/frontend-hackathon -f ./frontend/Dockerfile --progress tty .
	@echo "Build darchlabs/frontend-hackathon docker image done ✔︎"

compose-dev:
	@echo "[compose-dev]: Running docker compose dev mode..."
	@docker-compose -f infra/docker-compose.yml up sync postgres

compose: build-back
	@echo "[compose-dev]: Running docker compose dev mode..."
	@docker-compose -f infra/docker-compose.yml up --scale nginx=0

compose-stop:
	@echo "[compose-stop]: Stopping docker compose dev mode..."
	@docker-compose -f infra/docker-compose.yml down