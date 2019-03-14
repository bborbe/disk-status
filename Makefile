REGISTRY ?= docker.io
IMAGE ?= bborbe/disk-status
VERSION  ?= latest
VERSIONS = $(VERSION)
VERSIONS += $(shell git fetch --tags; git tag -l --points-at HEAD)

all: test

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u golang.org/x/tools/cmd/goimports

build:
	@tags=""; \
	for i in $(VERSIONS); do \
		tags="$$tags -t $(REGISTRY)/$(IMAGE):$$i"; \
	done; \
	echo "docker build --no-cache --rm=true $$tags ."; \
	docker build --no-cache --rm=true $$tags .

clean:
	@for i in $(VERSIONS); do \
		echo "docker rmi $(REGISTRY)/$(IMAGE):$$i"; \
		docker rmi $(REGISTRY)/$(IMAGE):$$i || true; \
	done

upload:
	@for i in $(VERSIONS); do \
		echo "docker push $(REGISTRY)/$(IMAGE):$$i"; \
		docker push $(REGISTRY)/$(IMAGE):$$i; \
	done

versions:
	@for i in $(VERSIONS); do echo $$i; done;

precommit: ensure format test check addlicense
	@echo "ready to commit"

ensure:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

format:
	@go get golang.org/x/tools/cmd/goimports
	@find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	@find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

test:
	go test -cover -race $(shell go list ./... | grep -v /vendor/)

check: lint vet errcheck

lint:
	@go get golang.org/x/lint/golint
	@golint -min_confidence 1 $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet $(shell go list ./... | grep -v /vendor/)

errcheck:
	@go get github.com/kisielk/errcheck
	@errcheck -ignore '(Close|Write|Fprint)' $(shell go list ./... | grep -v /vendor/)

addlicense:
	go get github.com/google/addlicense
	addlicense -c "Benjamin Borbe" -y 2019 -l bsd ./*.go ./disk/*.go
