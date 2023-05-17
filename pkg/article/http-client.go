package article

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const URL = "https://qiita.com/api/v2"

func request(method, url string, body io.Reader) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Failed to read environment variables.")
		return nil, err
	}

	token := os.Getenv("QIITA_API_TOKEN")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("Error: Request generation failed.")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: Failed to execute request.")
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Failed to read response.")
		return nil, err
	}
	return bodyBytes, nil
}
