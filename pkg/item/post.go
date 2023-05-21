package item

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Post(filename string) {
	url := URL + "/items"

	postItem, err := ConvertArticle(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonData, err := json.Marshal(postItem)
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyBytes, err := request("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var item Item

	err = json.Unmarshal([]byte(bodyBytes), &item)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	item.Format()

	printItem(item)
	// TODO: ファイル名を変更する
	// TODO: params を更新する
}
