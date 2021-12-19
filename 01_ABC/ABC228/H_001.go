package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)

	type node struct{ h, w int }
	a := make([]node, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i].h, &a[i].w)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].h < a[j].h
	})

	const N = 200002
	c := make([]int, N)
	for i := 1; i < n+1; i++ {
		c[i] = c[i-1] + a[i].w
	}

	s := make([]int, N)
	for i := 1; i < n+1; i++ {
		s[i] = s[i-1] + a[i].h*a[i].w
	}

	dp := make([]int, N)
	var X func(int) int
	X = func(i int) int {
		return c[i]
	}
	var Y func(int) int
	Y = func(i int) int {
		return dp[i] + s[i]
	}
	var slope func(int, int) float64
	slope = func(i, j int) float64 {
		return float64(Y(i)-Y(j)) / float64(X(i)-X(j))
	}

	t, h := 1, 1
	q := make([]int, N)
	for i := 1; i <= n; i++ {
		for h < t && slope(q[h], q[h+1]) <= float64(a[i].h) {
			h++
		}
		dp[i] = Y(q[h]) - X(q[h])*a[i].h + x + c[i]*a[i].h - s[i]
		for h < t && slope(q[t], q[t-1]) >= slope(i, q[t-1]) {
			t--
		}
		t++
		q[t] = i
	}
	fmt.Println(dp[n])
}
