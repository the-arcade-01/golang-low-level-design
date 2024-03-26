build:
	@go build -o ./bin main.go

run:build
	@./bin/main