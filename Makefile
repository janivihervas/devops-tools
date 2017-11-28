GO_FILES := `find . -type f -name '*.go' -not -path "./vendor/*"`
GO_PACKAGES := `go list ./...`
GO_TOOLS := github.com/golang/lint/golint \
			github.com/kisielk/errcheck \
			golang.org/x/tools/cmd/goimports

default: lint test build-foo build-bar

install:
	dep ensure

install-tools:
	go get -u $(GO_TOOLS)

lint:
	go vet ./... && golint -set_exit_status $(GO_PACKAGES) && errcheck -ignoretests ./...

test:
	go test -race -cover ./...

