.PHONY: dev run test cover tidy fmt pre-commit build-docker run-docker

.DEFAULT: help
help:
	@echo "make dev"
	@echo "	setup development environment"
	@echo "make run"
	@echo "	run app"
	@echo "make test"
	@echo "	run go test"
	@echo "make cover"
	@echo "	run go test with -cover"
	@echo "make tidy"
	@echo "	run go mod tidy"
	@echo "make fmt"
	@echo "	run gofumpt"
	@echo "make pre-commit"
	@echo "	run pre-commit hooks"
	@echo "make build-docker"
	@echo "	build docker container"
	@echo "make run-docker"
	@echo "	run docker container"

check-gofumpt:
ifeq (, $(shell which gofumpt))
	$(error "No gofumpt in $(PATH), gofumpt (https://pkg.go.dev/mvdan.cc/gofumpt) is required")
endif

check-pre-commit:
ifeq (, $(shell which pre-commit))
 	$(error "No pre-commit in $(PATH), pre-commit (https://pre-commit.com) is required")
endif

checks: check-gofumpt check-pre-commit

dev: checks
	pre-commit install

run:
	go build ./cmd/server.go && ./server

test:
	go test -v ./...

cover:
	go test -cover -v ./...

tidy:
	go mod tidy -compat=1.17

fmt: check-gofumpt
	gofumpt -l -w .

pre-commit: check-pre-commit
	pre-commit

build-docker:
	docker build -t go-grpc-helloworld .

run-docker:
	docker run --rm go-grpc-helloworld