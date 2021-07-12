package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	res := 0
	if a == b {
		res = c
	} else if b == c {
		res = a
	} else if c == a {
		res = b
	}
	fmt.Println(res)
}
