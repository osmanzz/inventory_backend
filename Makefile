#!/bin/bash

.PHONY: core
core:
	@go build

.PHONY: all
all: core