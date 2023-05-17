package main

import (
	"fmt"
	"os"

	"qiita-api/pkg/article"
)

const errMessage = "Error: The command has too many or too few arguments."

func main() {
	argsLen := len(os.Args)
	if !(argsLen >= 2 && argsLen <= 3) {
		fmt.Println(errMessage)
		return
	}

	method := os.Args[1]
	switch method {
	case "GET_ALL":
		if argsLen != 2 {
			fmt.Println(errMessage)
			return
		}
		article.GetAll()
	case "GET":
		if argsLen != 3 {
			fmt.Println(errMessage)
			return
		}
		article.Get(os.Args[2])
	case "POST":
		// FIXME: 記事の指定方法を修正したら引数を追加する
		// if argsLen != 3 {
		// 	fmt.Println(errMessage)
		// 	return
		// }
		article.Post()
	case "PATCH":
		if argsLen != 3 {
			fmt.Println(errMessage)
			return
		}
		// TODO: 更新処理
	case "DELETE":
		if argsLen != 3 {
			fmt.Println(errMessage)
			return
		}
		article.Delete(os.Args[2])
	case "SYNC_ALL":
		if argsLen != 2 {
			fmt.Println(errMessage)
			return
		}
		// TODO: 一括同期処理
	case "SYNC":
		if argsLen != 2 {
			fmt.Println(errMessage)
			return
		}
		// TODO: 同期処理
	default:
		fmt.Println("Error: An invalid command was entered.")
		return
	}

	fmt.Println("Process finished.")
}
