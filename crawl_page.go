package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}

	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if maxedOut := cfg.checkMaximumPages(); maxedOut {
		return
	}

	fmt.Printf("parsing current url: %s\n", rawCurrentURL)

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("unable to parse raw URL provided: %v", rawCurrentURL)
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		fmt.Printf("Hostname: %s of currentURL is not the same as host name (%s) of baseURL\n", currentURL.Hostname(), cfg.baseURL.Hostname())
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not normalize url: %s", rawCurrentURL)
		return
	}

	if isFirst := cfg.addPageVisit(normalizedURL); !isFirst {
		return
	}

	fmt.Printf("crawling... %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching html %v", err)
		return
	}

	fetchedURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting urls from html body: %v", err)
		return
	}

	for _, url := range fetchedURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
