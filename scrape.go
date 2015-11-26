package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
)

func main() {
	// request and parse the front page
	resp, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "athing"
		}
		return false
	}
	yellow := color.New(color.FgYellow).SprintFunc()

	// grab all articles and print them
	articles := scrape.FindAll(root, matcher)
	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, yellow(scrape.Text(article)), scrape.Attr(article, "href"))
	}
}
