package item

import (
	"encoding/json"
	"fmt"
)

func Get(itemID string) {
	url := URL + "/items/" + itemID
	bodyBytes, err := request("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var article Item

	json.Unmarshal([]byte(bodyBytes), &article)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	article.Format()

	printItem(article)
}
