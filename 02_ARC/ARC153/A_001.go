package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	x += 99999
	g := x % 10
	s := x / 10 % 10
	b := x / 100 % 10
	q := x / 1000 % 10
	w := x / 10000 % 10
	sw := x / 100000
	fmt.Printf("%d%d%d%d%d%d%d%d%d\n", sw, sw, w, q, b, b, s, g, s)
}
