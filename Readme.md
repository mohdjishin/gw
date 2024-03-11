
# API Gateway Testing Guide

This README provides instructions on how to test the API Gateway. The gateway is designed to route requests to the appropriate services based on the request path. Before testing, ensure the gateway and all backend services are running.

## Requirements

- API Gateway running on port 8080
- Backend services (e.g., UserService and ProductService) configured and running as per the gateway configuration

## Running the API Gateway

Make sure the API Gateway is running by executing the Go application:

```shell
go run gw/main.go
```

This starts the API Gateway on http://localhost:8080.

## Testing User Service Routing

If the UserService is configured to be accessible through /UserService/*, use the following curl command to test routing to the UserService:

```
curl http://localhost:8080/UserService/users
```

You should receive a JSON response from the UserService, indicating that the API Gateway has successfully routed the request.

## Testing Product Service Routing

If the ProductService is configured to be accessible through /ProductService/*, use the following curl command to test routing to the ProductService:

```
curl http://localhost:8080/ProductService/products
```

You should receive a JSON response from the ProductService, confirming that the API Gateway has successfully routed the request.

## Using Docker Compose to Manage the Project

Docker Compose allows you to manage multi-container Docker applications. To use Docker Compose to start and stop your project, follow these steps:

### Starting the Project with Docker Compose

1. **Ensure Docker Compose is installed**: Docker Compose should be installed on your system. For installation instructions, visit the [official Docker documentation](https://docs.docker.com/compose/install/).

2. **Create a `docker-compose.yml` file**: This file should define your services, networks, and volumes. Place this file at the root of your project directory.

3. **Start the services**: Navigate to the root of your project directory in a terminal and run the following command to start all services defined in your `docker-compose.yml` file:

```shell
docker-compose up -d
```

This command starts the services in detached mode, running them in the background.

### Stopping the Project with Docker Compose

To stop and remove all the running services defined in the `docker-compose.yml` file, run:

```shell
docker-compose down
```

This command stops all running containers and removes containers, networks, volumes, and images created by `docker-compose up`.

