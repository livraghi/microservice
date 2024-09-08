.DEFAULT_GOAL := all
.PHONY:

include mk/go.mk
include mk/dependencies.mk
include mk/test.mk

SUB_MODS = $(dir $(wildcard */Makefile))