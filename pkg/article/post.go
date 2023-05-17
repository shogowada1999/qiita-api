package article

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Post() {
	url := URL + "/items"
	// TODO: 投稿データを本文中から利用するように機能追加
	postData := Article{
		Body:    "# From GO",
		Private: false,
		Tags: []Tag{
			{
				Name:     "Ruby",
				Versions: []string{"0.0.1"},
			},
		},
		Title: "Example title",
		Tweet: false,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyBytes, err := request("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var article Article

	json.Unmarshal([]byte(bodyBytes), &article)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	article.Format()

	printArticle(article)
}
