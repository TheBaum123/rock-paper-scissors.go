build-run:
	@echo starting build process...
	@go build ./main.go
	@echo build successfull
	@echo
	@echo made by:
	@cat ./logo.txt
	@sleep 2
	@clear
	@echo launching program...
	@sleep 1
	@clear
	@./main

build:
	@echo starting build process...
	@go build ./main.go
	@echo build successfull
	@echo
	@echo made by:
	@cat ./logo.txt
	@sleep 2

BINARY_NAME=rock-paper-scissors-go

build-cross-platform:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

clean:
	@go clean
	@rm ${BINARY_NAME}-darwin
	@rm ${BINARY_NAME}-linux
	@rm ${BINARY_NAME}-windows