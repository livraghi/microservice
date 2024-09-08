.PHONY: clean

SUB_MODS ?=

GOPATH := $(shell go env GOPATH)

GO := $(shell command -v go 2> /dev/null || echo "/tmp/go-not-installed")
$(GO):
	$(error "GO cli not installed, please install it")

GO_FORMAT := $(shell command -v gofmt 2> /dev/null || echo "/tmp/gofmt-not-installed")
$(GO_FORMAT):
	$(error "GO cli not installed, please install it")

STATIC_CHECK := $(GOPATH)/bin/staticcheck
$(STATIC_CHECK):
	@go install honnef.co/go/tools/cmd/staticcheck@latest

GOSEC := $(GOPATH)/bin/gosec
$(GOSEC):
	@go install github.com/securego/gosec/cmd/gosec@latest

GOVULNCHECK := $(GOPATH)/bin/govulncheck
$(GOVULNCHECK):
	@go install golang.org/x/vuln/cmd/govulncheck@latest

./dist:
	@mkdir -p ./dist

clean:
	@if [ -n "$(SUB_MODS)" ]; then \
		for mod in $(SUB_MODS); do \
			$(MAKE) -C $$mod clean; \
		done; \
	fi

	@echo "Cleaning up.."
	@rm -rf ./vendor
	@rm -rf ./dist