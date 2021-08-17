package main

import (
	"bufio"
	"fmt"
	"os"
)

var p = int64(200_003)
var inv2 = (p + 1) / 2

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	out.Flush()

	sqrt := make(map[int64]int64)
	for i := int64(0); i < p; i++ {
		sqrt[i*i%p] = i
	}

	var n, m int64
	fmt.Scan(&n, &m)
	res := make([]int64, n)
	for i := int64(0); i < n; i++ {
		res[i] = (2 * p * i) + (i*i)%p
		fmt.Fprint(out, res[i])
		if i != n-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprint(out, "\n")

	var q, x, a, b int64
	fmt.Scan(&q)
	for i := int64(0); i < q; i++ {
		fmt.Fscan(in, &x)
		a, b = solve(x, sqrt)
		fmt.Fprintln(out, a+1, " ", b+1)
	}
}

func solve(s int64, sqrt map[int64]int64) (int64, int64) {
	a := int64(s >> 20)
	b := int64(s & ((1 << 20) - 1))
	c := sqrt[(2*b-a*a)%p]
	x, y := a+c, a-c
	x, y = x*inv2%p, y*inv2%p
	if x < y {
		return x, y
	}
	return y, x
}
