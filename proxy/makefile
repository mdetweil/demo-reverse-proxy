DEVCONTAINER=dev
DEVDOCKERFILENAME=dockerfile
APINAME=proxy

clean: ## Clean the directories
	rm -rf /bin /gen || true

dev: ## Build the container with the local volume attached for dev purposes
	@docker stop $(DEVCONTAINER) || true
	@docker build -t $(DEVCONTAINER) -f $(DEVDOCKERFILENAME) .
	@docker run -it --rm -p 8088:8088 --net host -v $(PWD):/app $(DEVCONTAINER)

build: clean ## build the server
	@go build -o bin/$(APINAME) ./proxy