.PHONY: bin/holiday

APP=bin/holiday

r:
	@go run -race cmd/main.go

clean:
	@go clean