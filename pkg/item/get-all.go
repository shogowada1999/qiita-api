package item

import (
	"encoding/json"
	"fmt"
)

func GetAll() {
	url := URL + "/authenticated_user/items"
	bodyBytes, err := request("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var items []Item

	err = json.Unmarshal([]byte(bodyBytes), &items)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	for i := range items {
		items[i].FormatDates()
		items[i].TruncateTitle()
	}

	for _, item := range items {
		printItem(item)
	}
}
