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
