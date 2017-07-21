package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is room to grow. Extend the slice
		z = x[:zlen]
	} else {
		// there is insufficient space, Allocate a new array
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
		fmt.Printf("resize, zlen: %d, zcap: %d\n", zlen, zcap)
	}
	z[len(x)] = y
	return z
}

func main() {
	var a []int
	for i := 0; i < 10; i++ {
		a = appendInt(a, i)
	}
	fmt.Println(a)
}
