/*
 * Compilation: go build -o outline main.go
 * Execution: ./fetch https://golang.org | ./outline
 * Dependences: net/html (go get golang.org/x/net/html)
 *
 * The program uses recursion over the HTML node tree to print the structure of
 * the tree in outline.
 * As it encounters each element, it pushes the element’s tag onto a stack,
 * then prints the stack.
 *
 * $ go build -o fetch ../../ch1/fetch
 * $ go build -o outline main.go
 * $ ./fetch https://golang.org | ./outline
 * [html]
 * [html head]
 * [html head meta]
 * [html head meta]
 * [html head meta]
 * [html head title]
 * [html head link]
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
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
