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

	var item Item

	json.Unmarshal([]byte(bodyBytes), &item)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	item.Format()

	printItem(item)
}
