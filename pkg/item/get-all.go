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

	var articles []Item

	json.Unmarshal([]byte(bodyBytes), &articles)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	for i := range articles {
		articles[i].FormatDates()
		articles[i].TruncateTitle()
	}

	for _, article := range articles {
		printItem(article)
	}
}
