TEST_DIR := ./service/tests/

all: up orderService
	
up:
	docker-compose up -d

down:
	docker-compose down

orderService:
	go run service/cmd/main.go

test:
	go test $(TEST_DIR)
