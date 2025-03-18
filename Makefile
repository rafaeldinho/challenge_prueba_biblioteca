install:
	 go mod tidy
run:
	go run main.go
	
test:
	 go test ./... --cover 

format:
	 go fmt ./...

testcoverage:
	 go test --coverprofile=coverage.out ./... && go tool cover -func=coverage.out