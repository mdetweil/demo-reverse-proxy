# Getting Started

## Via docker-compose

### Launch the containers
In a new terminal window:
 - `docker-compose up -d`

## Via docker

### Start the gRPC server
In a new terminal window:
- Change to the server directory: `cd server`
- Start the container via: `make dev` 
- In the container, build the project: `make build`
- In the container, run the server `./bin/demo`

### Start the proxy
In a new terminal window:
- Change to the proxy directory: `cd proxy`
- Start the container via: `make dev`
- In the container, build the project: `make build`
- In the container, run the proxy: `./bin/proxy`

## Make requests
From the host machine via curl or postman:
- `curl -X GET 'localhost:8088/calc/add?a=3&b=3'`
- `curl -X GET 'localhost:8088/calc/multiply?a=3&b=3'`
- `curl -X GET 'localhost:8088/world'`
