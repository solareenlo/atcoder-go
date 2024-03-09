package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var y, m, d int
	fmt.Scan(&y, &m, &d)
	if d+1 > b {
		d = 0
		m++
	}
	if m > a {
		m = 1
		y++
	}
	fmt.Println(y, m, d+1)
}
