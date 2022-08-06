package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

const UNSPLASH_ENDPOINT = "https://api.unsplash.com/photos/random"
const BLOG_ENDPOINT = "https://www.k-sato-0130.com/"

type Response struct {
	ID   string `json:"id"`
	URLS struct {
		Raw string `json:"raw"`
	}
}

// snippet only
var result Response

func httpClient() *http.Client {
	return &http.Client{}
}

// Memo
// Client: Set ogp image to the hosted server
// Server: Fetch a ramdom image from Unsplash API -> Fetch meta title from the client -> Take a screenshot with headless crome -> return the image
func FetchRandomImageURL() string {
	var ACCESS_KEY = os.Getenv("UNSPLASH_ACCESS_KEY")
	client := httpClient()

	req, err := http.NewRequest("GET", UNSPLASH_ENDPOINT, nil)

	req.Header.Add("Authorization", "Client-ID "+ACCESS_KEY)

	res, err := client.Do(req)

	if err != nil {
		var errMsg = fmt.Sprintf("Failed to fetch: %s", err)
		log.Fatal(errMsg)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return result.URLS.Raw
}

func FetchOpenGraphTitle(title string) string {
	client := httpClient()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blog/%s", BLOG_ENDPOINT, title), nil)

	res, err := client.Do(req)

	if err != nil {
		var errMsg = fmt.Sprintf("Failed to fetch: %s", err)
		log.Fatal(errMsg)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	attr, _ := doc.Find("meta[property='og:title']").Attr("content")

	return attr
}
