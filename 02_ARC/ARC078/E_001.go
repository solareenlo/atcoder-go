package main

import (
	"fmt"
	"os"
)

func PRINT(n int) {
	fmt.Println("!", n)
	os.Exit(0)
}

func query(n int) bool {
	if n == 0 {
		PRINT(1)
	}
	fmt.Println("?", n)
	var s string
	fmt.Scan(&s)
	return s == "Y"
}

func main() {
	l := 1000000000
	var r int
	if query(l) {
		for query(l - 1) {
			l /= 10
		}
		PRINT(l)
	}
	l /= 10
	for !query(l) {
		l /= 10
	}
	r = l*10 - 1
	for l < r-1 {
		mid := (l + r) >> 1
		if query(mid * 10) {
			r = mid
		} else {
			l = mid
		}
	}
	PRINT(r)
}
