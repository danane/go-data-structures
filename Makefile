all: check_go_fmt deps lint test

.PHONY: deps-init
deps-init:
	rm -f go.mod go.sum
	go mod init
	go mod tidy

.PHONY: update-deps
update-deps:
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: lint
lint:
	command -v golangci-lint || (cd /usr/local ; wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest)
	golangci-lint run --disable-all \
	--deadline=10m \
	--skip-files \.*_mock\.*\.go \
	-E errcheck \
	-E govet \
	-E unused \
	-E gocyclo \
	-E golint \
	-E varcheck \
	-E structcheck \
	-E maligned \
	-E ineffassign \
	-E interfacer \
	-E unconvert \
	-E goconst \
	-E gosimple \
	-E staticcheck \
	-E gosec

.PHONY: test
test:
	go test -v -cover -race ./...

.PHONY: check_go_fmt
check_go_fmt:
	@if [ -n "$$(gofmt -d $$(find . -name '*.go'))" ]; then \
		>&2 echo "The .go sources aren't formatted. Please format them with 'go fmt'."; \
		exit 1; \
	fi