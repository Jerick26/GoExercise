/******************************************************************************
 *  Compilation:  go build append.go
 *  Execution:    ./append
 *
 *  append a int to int slice with doubling capacity when full
 *
 *  % ./append
 *  0  cap=1	[0]
 *  1  cap=2	[0 1]
 *  2  cap=4	[0 1 2]
 *  3  cap=4	[0 1 2 3]
 *  4  cap=8	[0 1 2 3 4]
 *  5  cap=8	[0 1 2 3 4 5]
 *  6  cap=8	[0 1 2 3 4 5 6]
 *  7  cap=8	[0 1 2 3 4 5 6 7]
 *  8  cap=16	[0 1 2 3 4 5 6 7 8]
 *  9  cap=16	[0 1 2 3 4 5 6 7 8 9]
 *
 ******************************************************************************/

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
	}
	z[len(x)] = y
	return z
}

func main() {
	var a []int
	for i := 0; i < 10; i++ {
		a = appendInt(a, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(a), a)
	}
}
