build:
	@go build -o ./bin/goddit

run: build
	@./bin/goddit
