NAMESPACE = $(shell echo Edufund Technical Assestment - Backend - Vicktor)
TIMESTAMP = $(shell date +%FT%T%z)

# this make file will help to shortenest development syntax
dev:
	APP_ENV=local go run main.go