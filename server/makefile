DEVCONTAINER=dev
DEVDOCKERFILENAME=dockerfile
APINAME=demo

clean: ## Clean the directories
	rm -rf /bin /gen || true

gen: clean ## Generate the goa boilerplate
	@goa gen github.com/mdetweil/demo-reverse-proxy/server/design

dev: ## Build the container with the local volume attached for dev purposes
	@docker stop $(DEVCONTAINER) || true
	@docker build -t $(DEVCONTAINER) -f $(DEVDOCKERFILENAME) .
	@docker run -it --rm -p 8080:8080 --net host -v $(PWD):/app $(DEVCONTAINER)

build: gen ## build the server
	@go build -o bin/$(APINAME) ./cmd/calc