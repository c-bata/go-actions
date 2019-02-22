.PHONY: fmt
fmt: $(SOURCES) ## Formatting source codes.
	@goimports -w $^

.PHONY: lint
lint: ## Run golint and go vet.
	@golint -set_exit_status=1 $(PKGS)
	@go vet $(PKGS)
