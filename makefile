BINARY_NAME=rock-paper-scissors-go

build-run:
	@echo starting build process...
	@go build -o ${BINARY_NAME} ./main.go
	@echo build successfull
	@echo
	@echo made by:
	@cat ./logo.txt
	@sleep 2
	@clear
	@echo launching program...
	@sleep 1
	@clear
	@./${BINARY_NAME}

build:
	@echo starting build process...
	@go build -o ${BINARY_NAME} ./main.go
	@echo build successfull
	@echo
	@echo made by:
	@cat ./logo.txt
	@sleep 2

build-cross-platform:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go
	@echo build successfull

clean:
	@go clean
	@rm -f ${BINARY_NAME}-darwin
	@rm -f ${BINARY_NAME}-linux
	@rm -f ${BINARY_NAME}-windows
	@rm -f ${BINARY_NAME}
	@echo cleaned successfully
