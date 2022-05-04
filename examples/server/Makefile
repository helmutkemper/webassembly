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

.PHONY: build
## Executa o teste TestLocalDevOps sem habilitar o teste de caos, sem log de dados na saída padrão
build:
	@$(GO) mod tidy
	$(GO) build . server
	@if [ -f ./server.exe ]; then ./server.exe; fi
	@if [ -f ./server ]; then ./server; fi
