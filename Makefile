# Docker compose file name
DOCKER_COMPOSE_FILE = docker-compose.yml

# Run go application
run:
	@echo "Starting go application..."
	go run cmd/main.go

# Start the docker-compose services
start:
	@echo "Starting services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "Services started successfully!"

# Pause (stop) the containers without removing them
stop:
	@echo "Stopping Docker Compose services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) stop
	@echo "Services stopped!"

# Completely remove containers, volumes, and images
kill:
	@echo "Removing containers, volumes, and images..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes --rmi all
	@echo "All containers, volumes, and images removed!"

# Show logs of services in real time
logs:
	@echo "Displaying services' logs..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f