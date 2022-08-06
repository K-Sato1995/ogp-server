package main

import (
	"fmt"
	"ogp-server/api"
)

func main() {
	// fmt.Println(api.FetchImage())
	fmt.Println(api.FetchMeta("two-jp-baseball-super-stars"))
}
