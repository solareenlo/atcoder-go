package main

import "fmt"

func main() {
	var w, a, b int
	fmt.Scan(&w, &a, &b)

	res := 0
	if a <= b {
		if a+w >= b {
			res = 0
		} else {
			res = b - (a + w)
		}
	} else {
		if b+w >= a {
			res = 0
		} else {
			res = a - (b + w)
		}
	}
	fmt.Println(res)
}
