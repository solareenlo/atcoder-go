package main

import "fmt"

func main() {
	var n, m int
	var s, t string
	fmt.Scan(&n, &m, &s, &t)
	b := t[0:n]
	a := t[m-n : m]
	res := 2
	if b == s {
		res = 0
	}
	if a != s {
		res++
	}
	fmt.Println(res)
}
