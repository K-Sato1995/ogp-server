package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/chromedp/chromedp"
)

const UNSPLASH_ENDPOINT = "https://api.unsplash.com/photos/random"
const BLOG_ENDPOINT = "https://www.k-sato-0130.com/"

// func init() {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// (1) Share url => come to the edege => open up headless chrome and send back a pic
type Res struct {
	ID   string `json:"id"`
	URLS struct {
		Raw string `json:"raw"`
	}
}

var result Res

func httpClient() *http.Client {
	return &http.Client{}
}

func FetchRandomImageURL() string {
	ACCESS_KEY := os.Getenv("UNSPLASH_ACCESS_KEY")
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

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	return result.URLS.Raw
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// imgUrl := FetchRandomImageURL()

	bytes := Screenshot()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(bytes)
}

func Screenshot() []byte {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot(`https://pkg.go.dev/`, `img.Homepage-logo`, &buf)); err != nil {
		log.Fatal(err)
	}
	// if err := ioutil.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
	return buf

}
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}
