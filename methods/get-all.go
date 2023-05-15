package methods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const URL = "https://qiita.com/api/v2/authenticated_user/items"

type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (a *Article) FormatDates() {
	var err error
	a.CreatedAt, err = formatDateString(a.CreatedAt)
	if err != nil {
		fmt.Println("Error formatting CreatedAt:", err)
	}

	a.UpdatedAt, err = formatDateString(a.UpdatedAt)
	if err != nil {
		fmt.Println("Error formatting UpdatedAt:", err)
	}
}

func (a *Article) TruncateTitle() {
	if len(a.Title) > 50 {
		a.Title = a.Title[:50]
	}
}

func formatDateString(input string) (string, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02 15:04:05"), nil
}

func GetAll() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
		return
	}

	token := os.Getenv("QIITA_API_TOKEN")

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var articles []Article

	json.Unmarshal([]byte(body), &articles)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	for i := range articles {
		articles[i].FormatDates()
		articles[i].TruncateTitle()
	}

	for _, article := range articles {
		fmt.Printf("id        : %s\n", article.ID)
		fmt.Printf("created_at: %s\n", article.CreatedAt)
		fmt.Printf("updated_at: %s\n", article.UpdatedAt)
		fmt.Printf("title     : %s\n", article.Title)
		fmt.Printf("body      : %s\n", article.Body)
	}
}
