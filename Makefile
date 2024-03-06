IMAGE_NAME = finder
PORT = 8080

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run:
	docker run -p $(PORT):$(PORT) $(IMAGE_NAME)

# Stop the Docker container, remove it, and purge the image
stop:
	docker stop $$(docker ps -q --filter "ancestor=$(IMAGE_NAME)") || true
	docker rm $$(docker ps -aq --filter "ancestor=$(IMAGE_NAME)") || true
	docker rmi $(IMAGE_NAME) || true
