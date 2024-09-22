#!/bin/bash

# Pull the latest changes
git pull origin main

# Rebuild the Docker image
docker build -t my-go-app:latest .

# Stop and remove the old container
docker stop my-go-container
docker rm my-go-container

# Run a new container with the updated image
docker run -d -p 0.0.0.0:3002:3002 --name my-go-container my-go-app:latest

# Print the status of the new container
docker ps | grep my-go-container
