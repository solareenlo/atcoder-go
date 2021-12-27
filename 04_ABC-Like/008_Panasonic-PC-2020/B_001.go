package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	res := 0
	if h == 1 || w == 1 {
		res = 1
	} else if h%2 == 0 || w%2 == 0 {
		res = h * w / 2
	} else {
		res = h*w/2 + 1
	}

	fmt.Println(res)
}
