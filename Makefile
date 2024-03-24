.DEFAULT_GOAL := all
.PHONY: all clean check check-format format check-code-analysis-vet check-modules check-static-check check-security check-vulnerability compile test build update_dependencies vendor

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
NC='\033[0m'

FAIL="$(RED)fail$(NC)"
SUCCESS="$(GREEN)success$(NC)"
SEP="----------------------------------------\\n"


GOPATH := $(shell go env GOPATH)

GO := $(shell command -v go 2> /dev/null || echo "/tmp/go-not-installed")
$(GO):
	$(error "GO lang cli not installed, please install it")

GO_FORMAT := $(shell command -v gofmt 2> /dev/null || echo "/tmp/gofmt-not-installed")
$(GO_FORMAT):
	$(error "GO lang cli not installed, please install it")

STATIC_CHECK := $(GOPATH)/bin/staticcheck
$(STATIC_CHECK):
	@go install honnef.co/go/tools/cmd/staticcheck@latest

GOSEC := $(GOPATH)/bin/gosec
$(GOSEC):
	@go install github.com/securego/gosec/cmd/gosec@latest

GOVULNCHECK := $(GOPATH)/bin/govulncheck
$(GOVULNCHECK):
	@go install golang.org/x/vuln/cmd/govulncheck@latest

MODULES := $(sort $(dir $(shell find . -name go.mod)))

clean:
	@echo "Cleaning up... $(SUCCESS)"
	@rm -rf ./vendor

format: | $(GO_FORMAT)
	@$(eval OUTPUT = `$(GO_FORMAT) -d -e -l -s .`)
	@if [ -n "$(OUTPUT)" ]; then \
		echo "The following files are formatted:"; \
		echo "$(SEP)\n$(OUTPUT)\n$(SEP)"; \
		$(GO_FORMAT) -s -w .; \
	else \
		echo "Nothing to do"; \
	fi

check-format: | $(GO_FORMAT)
	@$(eval OUTPUT = `$(GO_FORMAT) -d -e -l -s .`)
	@if [ -n "$(OUTPUT)" ]; then \
		echo "Checking code style... $(FAIL)"; \
		echo "The following files should be formatted:"; \
		echo "$(SEP)\n$(OUTPUT)\n$(SEP)"; \
		exit 1; \
	else \
		echo "Checking code style... $(SUCCESS)"; \
	fi

check-code-analysis-vet: | $(GO)
	@$(GO) vet -c=5 $(foreach modeule, $(MODULES), $(modeule)...) && echo "Run code analysis... $(SUCCESS)" || (echo "Run code analysis... $(FAIL)" && exit 1)

check-modules: | $(GO)
	@BASE_PATH=$(PWD); \
	for module in $(MODULES); do \
  		OUTPUT="$$OUTPUT\nChecking module: $$module"; \
		cd $$module; \
		RESULT=`$(GO) mod verify 2>&1 || HAS_ERROR=true;` \
	    OUTPUT="$$OUTPUT\n$$RESULT";\
	    cd $$BASE_PATH; \
	done; \
	if [ -n "$$HAS_ERROR" ]; then \
		echo "Checking modules... $(FAIL)"; \
		echo "$(SEP)$$OUTPUT\n$(SEP)"; \
		exit 1; \
	else \
		echo "Checking modules... $(SUCCESS)"; \
	fi;

check-static-check: | $(STATIC_CHECK)
	@$(STATIC_CHECK) -f stylish -fail all $(foreach modeule, $(MODULES), $(modeule)...) && echo "Run static check... $(SUCCESS)" || (echo "Run static check... $(FAIL)" && exit 1)

check-security: | $(GOSEC)
	$(GOSEC) -tests -sort $(foreach modeule, $(MODULES), $(modeule)...) && echo "Run security check... $(SUCCESS)" || (echo "Run security check... $(FAIL)" && exit 1)

check-vulnerability: | $(GOVULNCHECK)
	@$(GOVULNCHECK) -test -show verbose $(foreach modeule, $(MODULES), $(modeule)...) && echo "Run vulnerability check... $(SUCCESS)" || (echo "Run vulnerability check... $(FAIL)" && exit 1)

compile:
test: compile
build: compile test
check: check-format check-code-analysis-vet check-modules check-static-check check-vulnerability # check-security
all: clean build check

update_dependencies:
	@BASE_PATH=$(PWD); \
	for module in $(MODULES); do \
  		echo "Update module dependencies: $$module"; \
		cd $$module; \
		go get -u ./...; \
		go mod tidy; \
		go mod verify; \
	    cd $$BASE_PATH; \
	done; \
	go work sync; \
	echo "Dependencies updated... $(SUCCESS)"

vendor:
	@go work vendor
