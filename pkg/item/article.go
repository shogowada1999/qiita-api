package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func extractParams(paramsText, target string) string {
	params := strings.Split(paramsText, "\n")

	value := ""
	for _, param := range params {
		trimmed := strings.TrimSpace(param)
		if strings.HasPrefix(trimmed, target) {
			split := strings.SplitN(trimmed, ":", 2)
			if len(split) == 2 {
				value = strings.TrimSpace(split[1])
			}
		}
	}
	return value
}

func parseBool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, fmt.Errorf("Invalid boolean representation")
}

func parseTags(tagsJSON string) ([]Tag, error) {
	var tags []Tag
	err := json.Unmarshal([]byte(tagsJSON), &tags)
	if err != nil {
		return tags, fmt.Errorf("Error: ")
	}
	return tags, nil
}

func divideParamsAndBody(content string) (string, string) {
	params := ""
	body := ""
	divider := "===\n"
	start := strings.Index(content, divider)
	end := strings.LastIndex(content, divider)
	if start != -1 && end != -1 && start < end {
		params = content[start+4 : end]
		body = content[end+4:]
		body = strings.TrimLeftFunc(body, func(r rune) bool {
			return r == '\n'
		})
	} else {
		fmt.Println("Could not find section enclosed by '---'")
	}
	return params, body
}

func generateItem(content string) (Item, error) {
	var item Item
	params, body := divideParamsAndBody(content)
	title := extractParams(params, "title")
	tags, errTags := parseTags(extractParams(params, "tags"))
	private, errPrivate := parseBool(extractParams(params, "private"))
	tweet, errTweet := parseBool(extractParams(params, "tweet"))
	if title == "" || errTags != nil || errPrivate != nil || errTweet != nil {
		return item, fmt.Errorf("Error:")
	}

	item.Title = title
	item.Body = body
	item.Tags = tags
	item.Private = private
	item.Tweet = tweet
	return item, nil
}

func ConvertArticle(filename string) (Item, error) {
	var item Item
	filePath := "articles/" + filename
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	content := string(data)

	item, err = generateItem(content)
	if err != nil {
		fmt.Println(err)
		return item, fmt.Errorf("Error")
	}
	return item, nil
}

func Test() {
	//
}

func GenerateArticleFile() error {
	now := time.Now()
	timeString := now.Format("2006-01-02-15-04-05")
	filename := fmt.Sprintf("%s.md", timeString)
	filepath := "articles/" + filename
	data, err := ioutil.ReadFile("assets/template.md")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("New article file is generated.")
	return nil
}
