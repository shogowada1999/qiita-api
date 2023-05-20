package item

import "fmt"

func Delete(itemID string) {
	url := URL + "/items/" + itemID
	_, err := request("DELETE", url, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
