package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages["mick"] = 32
	fmt.Printf("ages[mick] = %d\n", ages["mick"]) // "32"
	delete(ages, "mick")                          // remove element ages["alice"]
	ages["bob"]++
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Println("=== \ntest sort")
	//var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Println("===\ntest lookup")
	if _, ok := ages["John"]; !ok {
		fmt.Printf("John is not a key in ages. ages = %v\n", ages)
	}
	fmt.Println("===\ntest equal")
	eq := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println("two map is equal? answer:", eq)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
