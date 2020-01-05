# Getting Started

## Start the gRPC server
Change to the server directory: `cd server`
Start the container via: `make dev` 
In the container, build the project: `make build`
In the container, run the server `./bin/demo`

## Start the proxy
In a new terminal window:
Change to the proxy directory: `cd proxy`
Start the container via: `make dev`
In the container, build the project: `make build`
In the container, run the proxy: `./bin/proxy`

## Make requests
From the host machine via curl or postman:
GET localhost:8080/calc/add?a=3&b=3
GET localhost:8080/calc/multiply?a=3&b=3
GET localhost:8080/world
