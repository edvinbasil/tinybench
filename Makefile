#!/usr/bin/env make

# set APP (appname)
APP=tinybench

# set MAIN
MAIN=./cmd/main.go

# set GO binary (`GO=go-bin make` to override)
GO=go

# set SHELL
SHELL = /bin/sh

# set .PHONY
.PHONY: all build clean

# set SUFFIXES
.SUFFIXES: .go .py

all: $(APP) 
	./$(APP)

$(APP):
build:
	$(GO) build -o $(APP) $(MAIN)

clean:
	-rm $(APP)

