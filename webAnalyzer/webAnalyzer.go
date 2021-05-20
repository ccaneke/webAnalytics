package webAnalyzer

import (
	"strings"

	"golang.org/x/net/html"
)

type WebPageContent struct {
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
