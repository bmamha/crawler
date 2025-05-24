package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=======================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Printf("========================")
	pageCount := sortPages(pages)

	for _, page := range pageCount {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.url)
	}
}

type pageCount struct {
	url   string
	count int
}

func sortPages(pages map[string]int) []pageCount {
	pc := []pageCount{}
	for url, count := range pages {
		pc = append(pc, pageCount{url: url, count: count})
	}
	sort.Slice(pc, func(i, j int) bool { return pc[i].count > pc[j].count })

	return pc
}
