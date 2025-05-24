package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	maxPages           int
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	defer cfg.mu.Unlock()
	cfg.mu.Lock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL] += 1
		fmt.Printf("Adding count to %s in our cfg.pages map\n", normalizedURL)
		return false
	}

	cfg.pages[normalizedURL] = 1
	fmt.Printf("Added %s entry to our cfg.pages map\n", normalizedURL)
	return true
}

func (cfg *config) checkMaximumPages() (maxedOut bool) {
	defer cfg.mu.Unlock()
	cfg.mu.Lock()
	return (len(cfg.pages)) > cfg.maxPages
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		maxPages:           maxPages,
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
