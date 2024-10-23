# Docker compose file name
DOCKER_COMPOSE_FILE = docker-compose.yml

# .env file name
ENV = ".env.test"

# Run go application
run:
	@echo "Starting go application..."
	go run app/app.go

# Start the docker-compose services
start:
	@echo "Starting services..."
	docker-compose --env-file $(ENV) -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "Services started successfully!"

# Pause (stop) the containers without removing them
stop:
	@echo "Stopping Docker Compose services..."
	docker-compose --env-file $(ENV) -f $(DOCKER_COMPOSE_FILE) stop
	@echo "Services stopped!"

# Completely remove containers, volumes, and images
kill:
	@echo "Removing containers, volumes, and images..."
	docker-compose --env-file $(ENV) -f $(DOCKER_COMPOSE_FILE) down --volumes --rmi all
	@echo "All containers, volumes, and images removed!"

# Show logs of services in real time
logs:
	@echo "Displaying services' logs..."
	docker-compose --env-file $(ENV) -f $(DOCKER_COMPOSE_FILE) logs -f