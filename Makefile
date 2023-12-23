.PHONY: build install clean fmt lint

APP_NAME=gituser
BIN_PATH=$(USERPROFILE)\bin

fmt:
	gofmt -s -l .

lint: fmt
	golangci-lint run

build:
	go build -o $(APP_NAME).exe

install: build
	if not exist $(BIN_PATH) mkdir $(BIN_PATH)
	copy $(APP_NAME).exe $(BIN_PATH)

clean:
	if exist $(APP_NAME).exe del $(APP_NAME).exe

test-release:
	goreleaser release --snapshot --skip-publish --rm-dist

#ForWindows