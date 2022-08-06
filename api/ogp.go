package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

const END_POINT = "https://api.unsplash.com/photos/random"

func httpClient() *http.Client {
	return &http.Client{}
}

func FetchImage() string {
	var ACCESS_KEY = os.Getenv("UNSPLASH_ACCESS_KEY")
	client := httpClient()

	req, err := http.NewRequest("GET", END_POINT, nil)

	req.Header.Add("Authorization", "Client-ID "+ACCESS_KEY)

	res, err := client.Do(req)

	if err != nil {
		var errMsg = fmt.Sprintf("Failed to fetch: %s", err)
		log.Fatal(errMsg)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	sb := string(body)
	return sb
}
