package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 5005
	const M = 205

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := [N]int{}
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] += a[i-1]
	}

	t := [M]int{}
	sum := [N]int{}
	v := [M][N]int{}
	p := [M][N]int{}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var x int
			fmt.Fscan(in, &x)
			for ; t[j] > 0 && v[j][t[j]] <= x; t[j]-- {
				sum[p[j][t[j]]] -= v[j][t[j]]
				sum[p[j][t[j]-1]] += v[j][t[j]]
			}
			sum[p[j][t[j]]] -= x
			t[j]++
			v[j][t[j]] = x
			p[j][t[j]] = i
			sum[i] += x
		}
		s := 0
		for j := i; j > 0; j-- {
			s += sum[j]
			ans = max(ans, s-a[i]+a[j])
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
