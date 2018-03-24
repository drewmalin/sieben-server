APP=sieben-server
VERSION=1.0.0

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

DEPCMD=dep

PROJECT_ROOT=github.com/$(APP)

build: build_osx

build_osx:
				env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/osx/$(APP) $(PROJECT_ROOT)/cmd

clean:
				$(GOCLEAN)
				rm -rf bin/

deps:
				$(DEPCMD) ensure

test:
				$(GOTEST) -v ./...