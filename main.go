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
	h, err := getHTML(baseURL)
	if err != nil {
		fmt.Printf("Unable to fetch html: %v\n", err)
	}

	fmt.Println(h)
}
