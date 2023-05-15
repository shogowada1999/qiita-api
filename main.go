package main

import (
	"fmt"
	"os"

	"qiita-api/methods"
)

func main() {
	argsLen := len(os.Args)
	if !(argsLen >= 1 && argsLen <= 3) {
		fmt.Println("Error: The command has too many or too few arguments.")
		return
	}

	method := os.Args[1]
	switch method {
	case "GET_ALL":
		methods.GetAll()
	case "GET":
		//
	case "POST":
		//
	case "PATCH":
		//
	case "DELETE":
		//
	case "SYNC_ALL":
		//
	case "SYNC":
		//
	default:
		fmt.Println("Error: An invalid command was entered.")
		return
	}

	fmt.Println("Process finished.")
	return
}
