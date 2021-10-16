package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == b {
		fmt.Println(a * 2)
	} else if a > b {
		res := a
		if a-1 >= b {
			res += a - 1
		} else {
			res += b
		}
		fmt.Println(res)
	} else {
		res := b
		if b-1 >= a {
			res += b - 1
		} else {
			res += a
		}
		fmt.Println(res)
	}
}
