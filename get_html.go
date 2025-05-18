package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if res.StatusCode > 400 {
		return "", errors.New("response with error code 400 or above")
	}

	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("content-type not text/html")
	}

	resBody := res.Body
	defer res.Body.Close()

	b, err := io.ReadAll(resBody)
	if err != nil {
		return "", errors.New("can not read response body")
	}

	return string(b), nil
}
