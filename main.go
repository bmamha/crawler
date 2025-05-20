package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	if len(args) < 2 {
		fmt.Println("no website provided")

		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[1]
	pages := map[string]int{}
	crawlPage(baseURL, baseURL, pages)
	for k, v := range pages {
		fmt.Printf("%s: %d\n", k, v)
	}
}
