REGISTRY ?= docker.io
IMAGE ?= bborbe/disk-status
ifeq ($(VERSION),)
	VERSION := $(shell git fetch --tags; git describe --tags `git rev-list --tags --max-count=1`)
endif

all: test install

install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install cmd/disk-status/*.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install cmd/disk-status-server/*.go

test:
	go test -cover -race $(shell go list ./... | grep -v /vendor/)

vet:
	go tool vet .
	go tool vet --shadow .

lint:
	golint -min_confidence 1 ./...

errcheck:
	errcheck -ignore '(Close|Write)' ./...

check: lint vet errcheck

goimports:
	go get golang.org/x/tools/cmd/goimports

format: goimports
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

prepare:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/bborbe/docker_utils/bin/docker_remote_tag_exists

clean:
	docker rmi $(REGISTRY)/$(IMAGE)-build:$(VERSION)
	docker rmi $(REGISTRY)/$(IMAGE):$(VERSION)

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o disk-status-server ./go/src/github.com/$(IMAGE)/cmd/disk-status-server

build:
	docker build --build-arg VERSION=$(VERSION) --no-cache --rm=true -t $(REGISTRY)/$(IMAGE)-build:$(VERSION) -f ./Dockerfile.build .
	docker run -t $(REGISTRY)/$(IMAGE)-build:$(VERSION) /bin/true
	docker cp `docker ps -q -n=1 -f ancestor=$(REGISTRY)/$(IMAGE)-build:$(VERSION) -f status=exited`:/disk-status-server .
	docker rm `docker ps -q -n=1 -f ancestor=$(REGISTRY)/$(IMAGE)-build:$(VERSION) -f status=exited`
	docker build --no-cache --rm=true --tag=$(REGISTRY)/$(IMAGE):$(VERSION) -f Dockerfile.static .
	rm disk-status-server

upload:
	docker push $(REGISTRY)/$(IMAGE):$(VERSION)

version:
	@echo $(VERSION)

docker_remote_tag_exists:
	@go get github.com/bborbe/docker_utils/bin/docker_remote_tag_exists

trigger: docker_remote_tag_exists
	@exists=`docker_remote_tag_exists \
		-registry=${REGISTRY} \
		-repository="${IMAGE}" \
		-credentialsfromfile \
		-tag="${VERSION}" \
		-logtostderr \
		-v=0`; \
	trigger="build"; \
	if [ "$${exists}" = "true" ]; then \
		trigger="skip"; \
	fi; \
	echo $${trigger}

run:
	docker run \
	-p 9090:9090 \
	-e PORT=9090 \
	-e PATH=/volume \
	-v /tmp:/volume \
	$(REGISTRY)/bborbe/disk-status:$(VERSION) \
	-logtostderr \
	-v=0

