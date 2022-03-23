package main

import "fmt"

func upDate(x int) int {
	for x != (x & -x) {
		x -= x & -x
	}
	return x
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == b {
		fmt.Println(1)
		return
	}

	k := a ^ b
	k = upDate(k)
	a &= k - 1
	b &= k - 1
	b = upDate(b)
	if b != 0 {
		b = (b << 1) - 1
	}

	if b < a {
		fmt.Println((k << 1) - a - (a - 1 - b))
	} else {
		fmt.Println((k << 1) - a)
	}
}
