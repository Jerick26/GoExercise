/*
 * Compilation: go build -o findlinks1 main.go
 * Execution: ./fetch https://golang.org | ./findlinks1
 * Dependences: net/html (go get golang.org/x/net/html)
 *
 * The progam parses the standard input as HTML, extracts the links
 * using a recursive visit function, and prints each discovered link.
 *
 * $ go build -o fetch ../../ch1/fetch
 * $ go build -o findlinks1 main.go
 * $ ./fetch https://golang.org | ./findlinks1
 * /
 * /
 * #
 * /doc/
 * /pkg/
 * /project/
 * /help/
 * /blog/
 * http://play.golang.org/
 * #
 * #
 * //tour.golang.org/
 * https://golang.org/dl/
 * //blog.golang.org/
 * https://developers.google.com/site-policies#restrictions
 * /LICENSE
 * /doc/tos.html
 * http://www.google.com/intl/en/policies/privacy/
 *
 */

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
