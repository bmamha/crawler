package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("unable to parse raw URL provided: %v", rawBaseURL)
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("unable to parse baseURL provided")
	}

	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not normalize url: %s", rawCurrentURL)
		return
	}
	_, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL] += 1
		return
	}

	pages[normalizedURL] = 1
	fmt.Printf("crawling... %s ", rawCurrentURL)

	htmlBody, _ := getHTML(rawCurrentURL)

	fetchedURLs, err := getURLsFromHTML(htmlBody, normalizedURL)
	if err != nil {
		return
	}

	for _, url := range fetchedURLs {
		crawlPage(rawBaseURL, url, pages)
	}
}
