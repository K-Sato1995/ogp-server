package main

import (
	"fmt"
	"ogp-server/api"
)

func main() {
	fmt.Println(api.Screenshot())
	fmt.Println(api.FetchRandomImageURL())
}
