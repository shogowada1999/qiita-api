SUB_DIR := /

.PHONY: list
list:
	@go run main.go LIST

.PHONY: new
new:
	@go run main.go NEW $(SUB_DIR)

.PHONY: get-all
get-all:
	@go run main.go GET_ALL

.PHONY: get
get:
	@go run main.go GET ${ID}

# FIXME: 記事の指定方法を修正したら引数を追加する
.PHONY: post
post:
	@go run main.go POST

.PHONY: delete
delete:
	@go run main.go DELETE ${ID}

# FIXME: プロジェクト完了後削除
.PHONY: test
test:
	@go run main.go TEST