package main

import (
	"fmt"
	"ogp-server/api"
)

func main() {
	fmt.Println(api.FetchRandomImageURL())
	// fmt.Println(api.FetchOpenGraphTitle("two-jp-baseball-super-stars"))
}
