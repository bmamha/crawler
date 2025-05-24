package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	maxConcurrency = 3
	maxPages       = 10
)

func main() {
	args := os.Args
	fmt.Println(args)
	if len(args) < 2 {
		fmt.Println("no website provided")

		os.Exit(1)
	}

	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[1]

	if len(args) > 2 {
		c, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("trouble converting input for maxConcurrency: %v", err)
			os.Exit(1)
		}
		maxConcurrency = c
	}

	if len(args) > 3 {
		pages, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("trouble converting input for maxPages: %v", err)
			os.Exit(1)
		}
		maxPages = pages
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for k, v := range cfg.pages {
		fmt.Printf("%s: %d\n", k, v)
	}
}
