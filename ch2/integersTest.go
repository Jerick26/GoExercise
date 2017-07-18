package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // "1", "5"
		}
	}

	var a int8 = -33
	var b int8 = 33
	fmt.Printf("a = %08b, %[1]d\n", a)              // a = -0100001, -33
	fmt.Printf("b = %08b, %[1]d\n", b)              // b = 00100001, 33
	fmt.Printf("a << 1: %08b, %[1]d\n", int8(a<<1)) // a << 1: -1000010, -66
	fmt.Printf("b << 1: %08b, %[1]d\n", int8(b<<1)) // b << 1: 01000010, 66
	fmt.Printf("a << 2: %08b, %[1]d\n", int8(a<<2)) // a << 2: 01111100, 124
	fmt.Printf("b << 2: %08b, %[1]d\n", int8(b<<2)) // b << 2: -1111100, -124
	fmt.Printf("a * 4: %08b, %[1]d\n", int8(a*4))   // a * 4: 01111100, 124
	fmt.Printf("b * 4: %08b, %[1]d\n", int8(b*4))   // b * 4: -1111100, -124
	fmt.Printf("a << 3: %08b, %[1]d\n", int8(a<<3)) // a << 3: -0001000, -8
	fmt.Printf("b << 3: %08b, %[1]d\n", int8(b<<3)) // b << 3: 00001000, 8
	fmt.Printf("a * 8: %08b, %[1]d\n", int8(a*8))   // a * 8: -0001000, -8
	fmt.Printf("b * 8: %08b, %[1]d\n", int8(b*8))   // b * 8: 00001000, 8
	fmt.Printf("a >> 1: %08b, %[1]d\n", int8(a>>1)) // a >> 1: -0010001, -17
	fmt.Printf("b >> 1: %08b, %[1]d\n", int8(b>>1)) // b >> 1: 00010000, 16
	fmt.Printf("a >> 2: %08b, %[1]d\n", int8(a>>2)) // a >> 2: -0001001, -9
	fmt.Printf("b >> 2: %08b, %[1]d\n", int8(b>>2)) // b >> 2: 00001000, 8
	fmt.Printf("a / 4: %08b, %[1]d\n", int8(a/4))   // a / 4: -0001000, -8
	fmt.Printf("b / 4: %08b, %[1]d\n", int8(b/4))   // b / 4: 00001000, 8
}
