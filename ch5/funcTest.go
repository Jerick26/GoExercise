package main

import (
	"fmt"
	"math"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

//func findLinksLog(url string) ([]string, error) {
//	log.Printf("findLinks %s", url)
//	return findLinks(url)
//}

func f(i, j, k int, s, t string) { /* ... */ }
func add(x int, y int) int       { return x + y }
func sub(x, y int) (z int)       { z = x - y; return }
func first(x int, _ int) int     { return x }
func zero(int, int) int          { return 0 }

// Names are particularly valuable when a function retur ns multiple results
// of the same type
//func Size(rect image.Rectangle) (width, height int)     {}
//func Split(path string) (dir, file string)              {}
//func HourMinSec(t time.Time) (hour, minute, second int) {}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

func main() {
	fmt.Println(hypot(3, 4))  // "5"
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
	fmt.Printf("%.6f\n", math.Sin(30))
	// test multi-valued
	//log.Println(findLinks(url)) // rarely used in pro duction code
	//links, err := findLinks(url)
	//log.Println(links, err)
	// Function Values
	f := square
	fmt.Println(f(3)) // "9"
	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"
	//f = product           // compile error: can't assign f(int, int) int to f(int) int
	var f2 func(int) int
	//f2(3) // panic: call of nil function
	if f2 != nil {
		f2(3)
	}
	// parameterize functions
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	//doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	//word, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return 10, 20
}
