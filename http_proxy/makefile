DEVCONTAINER=dev
DEVDOCKERFILENAME=dockerfile
APINAME=http_proxy

dev: ## Build the container with the local volume attached for dev purposes
	@docker stop $(DEVCONTAINER) || true
	@docker build -t $(DEVCONTAINER) -f $(DEVDOCKERFILENAME) .
	@docker run -it --rm -p 8088:8088 --net host -v $(PWD):/app $(DEVCONTAINER)

build:  ## build the server
	@go build -o bin/$(APINAME) ./cmd/http_proxy