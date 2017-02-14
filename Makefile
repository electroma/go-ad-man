all: push

BUILDTAGS=

RELEASE?=0.1.1
PREFIX?=builditdigital/go-ad-man
GOOS?=linux

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
  COMMIT := git-$(shell git rev-parse --short HEAD)
endif

PKG=github.com/electroma/go-ad-man

build: clean
	GOOS=${GOOS} go build \
		-ldflags "-s -w -X ${PKG}/version.RELEASE=${RELEASE} -X ${PKG}/version.COMMIT=${COMMIT} -X ${PKG}/version.REPO=${REPO_INFO}" \
		-o go-ad-man ${PKG}

container: build package

package:
	docker build --pull -t $(PREFIX):$(RELEASE) .

push: package
	docker push $(PREFIX):$(RELEASE)
	docker tag $(PREFIX):$(RELEASE) $(PREFIX):latest
	docker push $(PREFIX):latest

fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' $(shell go list ${PKG}/... | grep -v vendor) | xargs -L 1 sh -c

lint:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"golint {{.Dir}}/..."{{end}}' $(shell go list ${PKG}/... | grep -v vendor) | xargs -L 1 sh -c

test: fmt lint vet
	@echo "+ $@"
	@go test -v -race -tags "$(BUILDTAGS) cgo" $(shell go list ${PKG}/... | grep -v vendor)

cover:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' $(shell go list ${PKG}/... | grep -v vendor) | xargs -L 1 sh -c
	gover
	goveralls -coverprofile=gover.coverprofile -service travis-ci -repotoken ${COVERALLS_TOKEN}

vet:
	@echo "+ $@"
	@go vet $(shell go list ${PKG}/... | grep -v vendor)

clean:
	rm -f go-ad-man
