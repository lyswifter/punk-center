SHELL=/usr/bin/env bash

.PHONY: clean
clean:
	rm punk-center

.PHONY: all
all:
	go build -o punk-center *.go