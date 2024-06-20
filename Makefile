SHELL := /bin/bash

all: build install

build:
	@echo " GO csmarkdoc"
	go mod tidy; \
	go build -o csmarkdoc -ldflags "-X main.Version=$(shell date +%y%m%d)"

install:
	@echo " MV csmarkdoc /usr/bin/csmarkdoc"
	@sudo mv csmarkdoc /usr/bin/csmarkdoc