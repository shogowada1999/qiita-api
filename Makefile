.PHONY: get-all
get-all:
	@go run main.go GET_ALL

.PHONY: get
get:
	@go run main.go GET ${ID}

.PHONY: delete
delete:
	@go run main.go DELETE ${ID}

