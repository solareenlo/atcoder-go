package main

import "fmt"

func main() {
	var l, h, n int
	fmt.Scan(&l, &h, &n)

	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		switch {
		case a < l:
			fmt.Println(l - a)
		case h < a:
			fmt.Println(-1)
		default:
			fmt.Println(0)
		}
	}
}
