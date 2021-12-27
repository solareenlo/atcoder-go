package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	tmp := 4 * a * b
	tmp2 := c - a - b
	tmp3 := tmp2 * tmp2

	ok := false
	if tmp2 > 0 && tmp < tmp3 {
		ok = true
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
