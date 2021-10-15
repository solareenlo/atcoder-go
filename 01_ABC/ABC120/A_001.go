package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c <= b/a {
		fmt.Println(c)
	} else {
		fmt.Println(b / a)
	}
}
