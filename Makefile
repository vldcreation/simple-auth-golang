NAMESPACE = $(shell echo Edufund Technical Assestment - Backend - Vicktor)
TIMESTAMP = $(shell date +%FT%T%z)
IS_IN_PROGRESS = $(shell echo "Is In Progress")
IS_DONE = $(shell echo "Is Done")

# this make file will help to shortenest development syntax
dev:
	APP_ENV=local go run main.go
docker-up:
	docker-compose up
docker-down:
	docker-compose down --remove-orphans --volumes
docker-build:
	@echo "docker-build ${IS_IN_PROGRESS}"
	docker-compose build
	@echo "docker-build finished ${IS_DONE}"
edufund-rebuild:
	@echo "edufund-rebuild ${IS_IN_PROGRESS}"
	docker-compose up -d --no-deps --build edufund-svc
	@echo "edufund-rebuild ${IS_DONE}"
edufund-restart:
	docker-compose restart edufund-svc

## test-unit: will test with standard testing
.PHONY: test-unit
test-unit:
	@echo "make test-unit ${IS_IN_PROGRESS}"
	@go clean -testcache ./...
	@go test \
		--race -count=1 -cpu=1 -parallel=1 -timeout=90s -failfast -vet= \
		-cover -covermode=atomic -coverprofile=./.coverage/unit.out \
		./internal/usecase/... \