package main

import "fmt"

var s, n int

func pw(x, y, t int) bool {
	for y > 0 {
		y--
		t /= x
	}
	return t > 0
}

func bs(t int) int {
	l := 1000
	r := n
	for l < r {
		z := (l + r + 1) >> 1
		if pw(z, t, n) {
			l = z
		} else {
			r = z - 1
		}
	}
	return l
}

func p(x int) {
	for t := n; t > 0; t /= x {
		if t%x > 1 {
			return
		}
	}
	s++
}

func P() {
	fmt.Scan(&n)
	s = 0
	for i := 2; i <= 1000; i++ {
		p(i)
	}
	for i := 1; ; i++ {
		x := bs(i)
		if x <= 1000 {
			break
		}
		for j := x; j > x-100 && j > 1000; j-- {
			p(j)
		}
	}
	fmt.Println(s)
}

func main() {
	var N int
	fmt.Scan(&N)
	for N > 0 {
		N--
		P()
	}
}
