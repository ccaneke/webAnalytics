package main

import (
	"fmt"
	"net/http"

	"github.com/ccaneke/webAnalytics/webAnalyzer"
)

func main() {
	const url = "http://example.com"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var pageInformation webAnalyzer.PageInformation = webAnalyzer.GetPageInformation(resp.Body)

	fmt.Println(pageInformation)
}
