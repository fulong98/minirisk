.PHONY: build run test clean migrate-up migrate-down docker-build docker-up docker-down

# Development
build:
	cd backend && go build -o ../bin/minirisk-api

run: build
	./bin/minirisk-api

test:
	cd backend && go test ./...

clean:
	rm -rf bin/
	rm -rf logs/

# Database
migrate-up:
	docker-compose exec mysql mysql -u minirisk -pminirisk_password minirisk < database/migrations/001_initial_schema.sql

migrate-down:
	docker-compose exec mysql mysql -u minirisk -pminirisk_password minirisk -e "DROP TABLE IF EXISTS positions, market_data, margins;"

# Docker
docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Development environment setup
setup: docker-build docker-up migrate-up
	@echo "Development environment is ready!"

# Frontend development
frontend-install:
	cd frontend && npm install

frontend-dev:
	cd frontend && npm start

frontend-build:
	cd frontend && npm run build

# Backend development
backend-install:
	cd backend && go mod download

backend-dev:
	cd backend && go run main.go

# Help
help:
	@echo "Available commands:"
	@echo "  build          - Build the backend application"
	@echo "  run            - Run the backend application"
	@echo "  test           - Run tests"
	@echo "  clean          - Clean build artifacts"
	@echo "  migrate-up     - Apply database migrations"
	@echo "  migrate-down   - Rollback database migrations"
	@echo "  docker-build   - Build Docker images"
	@echo "  docker-up      - Start Docker containers"
	@echo "  docker-down    - Stop Docker containers"
	@echo "  setup          - Set up development environment"
	@echo "  frontend-install - Install frontend dependencies"
	@echo "  frontend-dev   - Start frontend development server"
	@echo "  frontend-build - Build frontend for production"
	@echo "  backend-install - Install backend dependencies"
	@echo "  backend-dev    - Start backend development server" 