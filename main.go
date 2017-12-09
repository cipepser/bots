package main

import (
	"fmt"

	"./facebook"
)

func main() {
	URL = "https://graph.facebook.com/v2.11/<user name>/feed"

	f, err := facebook.GetFeed(URL)
	if err != nil {
		panic(err)
	}

	fmt.Println(f)

}
