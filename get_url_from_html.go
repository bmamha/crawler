package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlbody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlbody)
	n, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	links := []string{}
	links, err = extractLinks(n, links, baseURL)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func extractLinks(n *html.Node, links []string, baseURL *url.URL) ([]string, error) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				if attr.Val == "" {
					continue
				}

				hrefURL, err := url.Parse(attr.Val)
				if err != nil {
					fmt.Println("unable to parse href")
					continue
				}

				resolvedURL := baseURL.ResolveReference(hrefURL)
				fmt.Println(resolvedURL)
				links = append(links, resolvedURL.String())

			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links, _ = extractLinks(c, links, baseURL)
	}
	return links, nil
}
