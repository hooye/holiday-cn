.PHONY: bin/holiday

APP=bin/holiday

r:
	@go run -race cmd/* run

t: 
	@go run -race cmd/* run -t

clean:
	@go clean