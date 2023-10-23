package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e18)

var A, ans [200000]int

func rec(a, b, l, r int) {
	if b <= a {
		return
	}
	res := INF
	mid := (a + b) / 2
	idx := l
	for i := l; i < r; i++ {
		tmp := A[i] + (mid-i)*(mid-i)
		if tmp < res {
			res = tmp
			idx = i
		}
	}
	ans[mid] = res
	rec(a, mid, l, idx+1)
	rec(mid+1, b, idx, r)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	rec(0, N, 0, N)
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
