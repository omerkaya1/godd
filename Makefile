# Tabs are the main pain when it commes to Makefile debugging, so after you're done with the changes
# in the Makefile, run 'cat -e -t -v Makefile' to see where tabs were substituted with spaces
# ^I indicates \t and $ indicates \r or \n. Both are vital and everything else before and after may
# cause make errors.

export GO111MODULE=on
BIN = $(CURDIR)/build
# export PATH=$PATH:$GOPATH/bin

.PHONY: setup
setup: ## Install all the build and lint dependencies
	cd $(BIN)
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u golang.org/x/tools/cmd/cover
	go get -u golang.org/x/tools/cmd/goimports
	cd ../

.PHONY: mod
mod: ## Runs mod
	go mod vendor
	go mod verify
	go mod tidy

.PHONY: test
test: ## Runs all the tests
	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: cover
cover: test ## Runs all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: fmt
fmt: setup ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: lint
lint: setup ## Runs all the linters
	$(BIN)/golangci-lint run --disable-all \
		--enable=staticcheck \
		--enable=gosimple \
		--enable=gofmt \
		--enable=golint \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=10m \
		./...

.PHONY: build
build: ## Builds the project
	go build -o $(GOPATH)build/godd

.PHONY: clean
clean: ## Remove temporary files
	go clean
	rm -rf $(BIN)
	rm testdata/output_file.txt

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
