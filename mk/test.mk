.PHONY: test unit_tests
DEFAULT_GOAL := test

SUB_MODS ?=

unit_tests: ./dist | $(GO)
	@if [ -n "$(SUB_MODS)" ]; then \
		for mod in $(SUB_MODS); do \
			$(MAKE) -C $$mod unit_tests; \
		done; \
	fi

	@$(GO) test -count=1 -cover -covermode=count -outputdir=./dist/ .
	$(MAKE) -C test/unit-tests

test: unit_tests
	@echo "All tests passed"