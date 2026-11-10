BINARY=server.exe
up:
	@echo "starting... "
	docker-compose up --build -d --remove-orphans

down:
	docker-compose down

build:
	go build -o $(BINARY) ./api

start:
	./$(BINARY)
  
restart: build start
	