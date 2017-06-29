GO_FILES = $(shell find . -name '*.go' -not -path './vendor/*')
GO_PACKAGES = $(shell go list ./...  | grep -v /vendor/)
GO_TOOLS := github.com/golang/lint/golint \
			github.com/kisielk/errcheck \
			golang.org/x/tools/cmd/goimports

install-vendor:
	git submodule update --init

install-tools:
	go get -u $(GO_TOOLS)

build-foo:
	go install github.com/janivihervas/devops-tools/cmd/foo

build-bar:
	go install github.com/janivihervas/devops-tools/cmd/bar

lint:
	go vet -v $(GO_PACKAGES) && golint -set_exit_status $(GO_PACKAGES) && errcheck -ignoretests $(GO_PACKAGES)

test:
	go test -race $(GO_PACKAGES)

