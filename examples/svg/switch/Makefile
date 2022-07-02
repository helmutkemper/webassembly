GO ?= go

.PHONY: help
## Este comando de ajuda
help:
	@printf "Opções de comandos\n\n"
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "make %-30s ## %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: buildandrun
## build this example and run local server
buildandrun:
	@$(GO) mod tidy
	@GOARCH=wasm GOOS=js $(GO) build -o main.wasm
	@$(MAKE) -C ../../server build


.PHONY: build
## build this example
build:
	$(GO) mod tidy
	GOARCH=wasm GOOS=js $(GO) build -o main.wasm


.PHONY: server
## run local server
server:
	@$(GO) mod tidy
	@$(MAKE) -C ../../server build