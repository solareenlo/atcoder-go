package main

import "fmt"

func main() {
	var b, c int
	fmt.Scan(&b, &c)

	var res int
	if b == 0 {
		res = c
	} else if c == 0 {
		res = 1
	} else if c >= 2*abs(b) {
		tmp := 0
		if b > 0 {
			tmp = 1
		}
		res = c + 2*abs(b) - tmp
	} else if c == 1 {
		res = 2
	} else {
		res = 2*c - 1
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
