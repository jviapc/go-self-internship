package main

import (
	"fmt"
	"html"
	"regexp"
	. "wiki/parser/infrastructure/httpclient"
)

func main() {
	// Выполняем HTTP-запрос к странице с фамилиями (пример страницы)
	urlPrefix := "https://en.wikipedia.org"

	url := "/wiki/Category:Given_names?from=A"

	httpClient := NewClientLogger(new(Client))

	for len(url) > 0 {
		body, _ := httpClient.Get(urlPrefix + url)

		body = convertHtmlEntities(body)

		url = findNextPageUrl(body)

		fmt.Println(findNames(body))
	}
}

func convertHtmlEntities(val []byte) []byte {
	return []byte(html.UnescapeString(string(val)))
}

func findNames(html []byte) []string {
	nameRegexp := regexp.MustCompile(
		`<li>\s*(?:<span class="redirect-in-category">)?<a\s*href="/wiki[^:>]+?>(.*?)(?:\s*\((?:given\s*)?name\))?</a>(?:</span>)?</li>`,
	)

	names := nameRegexp.FindAllSubmatch(html, -1)

	var output []string

	for _, element := range names {
		output = append(output, string(element[1]))
	}

	return output
}

func findNextPageUrl(html []byte) string {
	nextPageRegexp := regexp.MustCompile(
		`<a href="([^"]+?)" title="[^"]+?">next page</a>`,
	)

	nextPage := nextPageRegexp.FindSubmatch(html)

	if len(nextPage) < 2 {
		return ""
	}

	return string(nextPage[1])
}
