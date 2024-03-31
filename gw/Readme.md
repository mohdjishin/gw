# API Gateway Testing Guide

This README provides instructions on how to test the API Gateway. The gateway is designed to route requests to the appropriate services based on the request path. Before testing, ensure the gateway and all backend services are running.

## Requirements

- API Gateway running on port 8080
- Backend services (e.g., UserService and ProductService) configured and running as per the gateway configuration

## Running the API Gateway

Make sure the API Gateway is running by executing the Go application:

```shell
go run *.go
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

