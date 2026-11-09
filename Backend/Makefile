include .env
up:
	@echo "starting... "
	docker-compose up --build -d --remove-orphans

down:
	docker-compose down

build:
	go build -o $(BINARY) ./api

start:
	./ $(BINARY)
  
restart: down up
	