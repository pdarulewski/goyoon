lint:
  golangci-lint run
  sqlc diff
  sqlc vet
  sqlc verify

test:
  gotestsum --format dots

generate:
  rm -rf ./internal/repository
  sqlc generate
