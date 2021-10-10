package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	res := ""
	if n == 0 {
		res = "0"
	}
	for n != 0 {
		if n%2 == 0 {
			res = "0" + res
		} else {
			n--
			res = "1" + res
		}
		n /= -2
	}
	fmt.Println(res)
}
