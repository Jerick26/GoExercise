/*
 * Compilation: Go build mapKey.go
 * Execution: ./mapKey
 *
 * the program uses a map to record the number of times Add has been called
 * with a given list of strings. It uses fmt.Sprintf to convert a slice of str
 * ings into a single str ing that is a suit able map key, quoting each slice
 * element wit h %q to record str ing bound aries faithfully
 *
 * $ ./mapKey
 * map[["rice" "tomato"]:3 ["tomato"]:3 ["tomato" "apple"]:1]
 * Count ["rice", "tomato"]:  3
 *
 */
package main

import "fmt"

func main() {
	var m = make(map[string]int)
	a := []string{"rice", "tomato", "apple", "banana", "orange",
		"rice", "tomato", "apple", "peach", "pear"}
	Add(m, a[:2])
	Add(m, a[:2])
	Add(m, a[5:7])
	Add(m, a[1:2])
	Add(m, a[6:7])
	Add(m, a[1:3])
	Add(m, a[6:7])
	fmt.Println(m)
	fmt.Println(`Count ["rice", "tomato"]: `, Count(m, a[:2]))
}

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(m map[string]int, list []string) { m[k(list)]++ }

func Count(m map[string]int, list []string) int { return m[k(list)] }
