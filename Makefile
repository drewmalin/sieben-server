APP=sieben-server
VERSION=1.0.0

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt

DEPCMD=dep

PROJECT_ROOT=github.com/$(APP)

build: clean fmt build_osx build_linux build_win test

build_osx:
		env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/osx/$(APP) $(PROJECT_ROOT)/cmd

build_linux:
		env GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/linux/$(APP) $(PROJECT_ROOT)/cmd

build_win:
		env GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/win64/$(APP) $(PROJECT_ROOT)/cmd

clean:
		$(GOCLEAN)
		rm -rf bin/

deps:
		$(DEPCMD) ensure

fmt:
		$(GOFMT) ./...

test:
		$(GOTEST) -v ./...
