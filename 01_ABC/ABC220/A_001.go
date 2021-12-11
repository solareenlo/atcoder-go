package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	for i := c; i <= 1000; i += c {
		if a <= i && i <= b {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
