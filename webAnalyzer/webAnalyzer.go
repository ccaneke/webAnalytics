package webAnalyzer

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// A PageInformation consists of information about a web page
type PageInformation struct {
	HtmlVersion          string
	PageTitle            string
	NumH1                int
	NumH2                int
	NumH3                int
	NumH4                int
	NumAnchors           int
	NumLinks             int
	NumInaccessibleLinks int
	LoginForm            bool
}

// Contains verifies that a web page contains a login form
func Contains(n *html.Node) bool {
	var s []html.Attribute = n.Attr

	for _, v := range s {
		if v.Key == "type" && strings.ToLower(v.Val) == "password" {
			return true
		}

	}
	return false
}

// GetHTMLVersion gets the HTML version of a web page
func GetHTMLVersion(r io.Reader) string {
	var version string
	var scanner = bufio.NewScanner(r)
	var line string

	for i := 0; scanner.Scan() && i < 1; i++ {
		line = scanner.Text()
	}

	if !strings.Contains(strings.ToLower(line), "public") {
		version = "HTML5"
	} else {
		version = "HTML4"
	}

	return version
}

// GetPageInformation gets information about the contents of a web page
func GetPageInformation(r io.Reader) (PageInformation, error) {
	b, err := io.ReadAll(r)
	var page = PageInformation{}
	if err != nil {
		return page, err
	}

	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		return page, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if (*n).Type == html.ElementNode && (*n).Data == "link" {
			page.NumLinks += 1
		}

		if (*n).Type == html.ElementNode && (*n).Data == "a" {
			page.NumAnchors += 1

			url := (*n).Attr[0].Val
			_, err := http.Get(url)

			if err != nil {
				page.NumInaccessibleLinks += 1
			}
		}

		if (*n).Type == html.ElementNode && (*n).Data == "h1" {
			page.NumH1 += 1
		}

		if (*n).Type == html.ElementNode && (*n).Data == "h2" {
			page.NumH2 += 1
		}

		if (*n).Type == html.ElementNode && (*n).Data == "h3" {
			page.NumH3 += 1
		}

		if (*n).Type == html.ElementNode && (*n).Data == "h4" {
			page.NumH4 += 1
		}

		if Contains(n) {
			page.LoginForm = true
		}

		if (*n).Parent != nil && (*n).Parent.Type == html.ElementNode && (*n).Parent.Data == "title" && (*n).Type == html.TextNode {
			page.PageTitle = (*n).Data
		}

		for c := (*n).FirstChild; c != nil; c = (*c).NextSibling {
			f(c)
		}
	}
	f(doc)

	page.HtmlVersion = GetHTMLVersion(bytes.NewReader(b))

	return page, nil
}
