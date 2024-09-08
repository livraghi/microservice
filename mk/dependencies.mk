.PHONY: show_sub_mods

SUB_MODS ?=

show_sub_mods:
	@echo $(SUB_MODS)
	@if [ -n "$(SUB_MODS)" ]; then \
		for mod in $(SUB_MODS); do \
			$(MAKE) -C $$mod show_sub_mods; \
		done; \
	fi

update_dependencies: | $(GO)
	@$(GO) get -u .
	@$(GO) mod tidy

	@if [ -n "$(SUB_MODS)" ]; then \
		for mod in $(SUB_MODS); do \
			$(MAKE) -C $$mod update_dependencies; \
		done; \
	fi
	@echo "Dependencies updated"