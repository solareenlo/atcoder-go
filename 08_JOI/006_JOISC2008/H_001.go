package main

import "fmt"

var M, n, k int
var NA bool

func solve(a, b, c, d int) {
	if (b + d) > M {
		return
	}
	if n > k {
		return
	}
	solve(a, b, a+c, b+d)
	n++
	if n == k {
		fmt.Println(a+c, b+d)
		NA = false
		return
	}
	solve(a+c, b+d, c, d)
}

func main() {
	NA = true
	fmt.Scan(&M, &k)
	solve(0, 1, 1, 1)
	if NA {
		fmt.Println(-1)
	}
}
