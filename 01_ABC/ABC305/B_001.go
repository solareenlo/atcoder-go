package main

import "fmt"

func main() {
	var x [7]int
	x[0] = 0
	x[1] = 3
	x[2] = x[1] + 1
	x[3] = x[2] + 4
	x[4] = x[3] + 1
	x[5] = x[4] + 5
	x[6] = x[5] + 9
	var p, q string
	fmt.Scan(&p, &q)
	a := p[0] - 'A'
	b := q[0] - 'A'
	if a > b {
		a, b = b, a
	}
	fmt.Println(x[b] - x[a])
}
