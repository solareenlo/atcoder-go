package main

import (
	"bufio"
	"fmt"
	"os"
)

func get(n, x int, A, B int) int {
	if x > n {
		return 0
	}
	m := n/x - 1
	return ((A+A-(m-1)*B)*m/2)*x + (A-m*B)*(n%x)
}

func solve(n int, A, B int) int {
	l, r := 1, n
	var ans int
	for l <= r {
		md := (l + r) / 2
		if get(n, md, A, B) >= get(n, md+1, A, B) {
			ans, r = md, md-1
		} else {
			l = md + 1
		}
	}
	return get(n, ans, A, B)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	var A, B int
	fmt.Fscan(in, &n, &A, &B, &m)
	x := make([]int, m+2)
	x[m+1] = n + 1
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &x[i])
	}
	var ans int
	for i := 0; i <= m; i++ {
		ans += solve(x[i+1]-x[i], A, B)
	}
	fmt.Println(ans)
}
