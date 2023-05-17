package article

import (
	"fmt"
	"time"
)

type Tag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}
type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Tags      []Tag  `json:"tags"`
	Private   bool   `json:"private"`
	Tweet     bool   `json:"tweet"`
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

func (a *Article) Format() {
	a.FormatDates()
	a.TruncateTitle()
}

func formatDateString(input string) (string, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02 15:04:05"), nil
}

func printArticle(a Article) {
	fmt.Printf("id        : %s\n", a.ID)
	fmt.Printf("private   : %t\n", a.Private)
	fmt.Printf("created_at: %s\n", a.CreatedAt)
	fmt.Printf("updated_at: %s\n", a.UpdatedAt)
	fmt.Printf("tags      : %s\n", a.Tags)
	fmt.Printf("title     : %s\n", a.Title)
	fmt.Printf("body      : %s\n\n", a.Body)
}
