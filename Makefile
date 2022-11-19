NAMESPACE = $(shell echo Edufund Technical Assestment - Backend - Vicktor)
TIMESTAMP = $(shell date +%FT%T%z)
IS_IN_PROGRESS = $(shell echo "Is In Progress")

# this make file will help to shortenest development syntax
dev:
	APP_ENV=local go run main.go

## test-unit: will test with standard testing
.PHONY: test-unit
test-unit:
	@echo "make test-unit ${IS_IN_PROGRESS}"
	@go clean -testcache ./...
	@go test \
		--race -count=1 -cpu=1 -parallel=1 -timeout=90s -failfast -vet= \
		-cover -covermode=atomic -coverprofile=./.coverage/unit.out \
		./internal/usecase/... \