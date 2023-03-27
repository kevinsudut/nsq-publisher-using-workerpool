# go build command
gobuild:
	@go build -v -o bin/main cmd/main.go

# go run command
gorun:
	make gobuild
	@./bin/main

# docker compose up command
composeup:
	@docker compose up -d

# docker compose down command
composedown:
	@docker compose down
