.PHONY: build run dev stop logs

# Build all containers
build:
	docker-compose build

# Run all containers in background
run:
	docker-compose up -d

dev:
	docker-compose up --build

# Stop and remove containers
stop:
	docker-compose down

# Follow logs
logs:
	docker-compose logs -f


test:
	curl -X POST http://localhost:8081/register -H "Content-Type: application/json" -d '{"name":"kaleab","email":"k@a.com","password":"123"}'
	curl -X POST http://localhost:8081/login -H "Content-Type: application/json" -d '{"email":"k@a.com","password":"123"}'

user-service:
	cd user-service && go run main.go

book-service:
	cd book-service && go run main.go
