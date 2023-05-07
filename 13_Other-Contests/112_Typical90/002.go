package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	rec(n, 0, "")
}

func rec(n, c int, s string) {
	if n == 0 {
		fmt.Println(s)
		return
	}
	n--
	if c+1 <= n {
		rec(n, c+1, s+"(")
	}
	if c != 0 {
		rec(n, c-1, s+")")
	}
}
