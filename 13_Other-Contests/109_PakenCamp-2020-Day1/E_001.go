package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x == y && x != 0 {
		fmt.Println(-1)
	} else {
		var a, b int
		if x == y {
			a = 1
			b = 1
		} else if x < y {
			b = y
			a = b + x
			if x == 0 {
				a *= 2
			}
		} else {
			a = x
			b = a + y
			if y == 0 {
				b *= 2
			}
		}
		fmt.Println(a, b)
	}
}
