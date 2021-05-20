package webAnalyzer

import (
	"bufio"
	"io"
	"strings"

	"golang.org/x/net/html"
)

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

func Contains(n *html.Node) bool {
	var s []html.Attribute = n.Attr

	for _, v := range s {
		if v.Key == "input" && strings.ToLower(v.Val) == "password" {
			return true
		}

	}
	return false
}

func GetHTMLVersion(r io.Reader) string {
	var version string
	var scanner = bufio.NewScanner(r)
	var line string

	for i := 0; scanner.Scan() && i < 1; i++ {
		line = scanner.Text()
	}

	if strings.Contains(line, "!DOCTYPE html") {
		version = "HTML5"
	} else {
		version = "HTML4"
	}

	return version
}
