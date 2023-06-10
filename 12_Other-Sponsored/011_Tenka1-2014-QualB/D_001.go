package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N)
	for j := 0; j < N; j++ {
		fmt.Fscan(in, &A[j])
	}
	var H, W int
	fmt.Fscan(in, &H, &W)
	B := make([][]int, H)
	for i := range B {
		B[i] = make([]int, W)
	}
	l := make([]int, N)
	r := make([]int, N)
	u := make([]int, N)
	d := make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = INF
		r[i] = -INF
		u[i] = INF
		d[i] = -INF
	}
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			var i int
			fmt.Fscan(in, &i)
			B[y][x] = i
			l[i] = min(l[i], x)
			r[i] = max(r[i], x)
			u[i] = min(u[i], y)
			d[i] = max(d[i], y)
		}
	}
	v := make([]int, N)
	for i := 0; i < N; i++ {
		if l[i] == INF {
			continue
		}
		for y := u[i]; y < d[i]+1; y++ {
			for x := l[i]; x < r[i]+1; x++ {
				if B[y][x] != i {
					v[B[y][x]] |= (1 << i)
				}
			}
		}
	}
	w := make([]int, 1<<N)
	for S := 0; S < 1<<N; S++ {
		for i := 0; i < N; i++ {
			if (S & (1 << i)) != 0 {
				w[S] |= v[i]
			}
		}
	}
	dp := make([]bool, 1<<N)
	dp[0] = true
	sort.Ints(A)
	ok := make([]bool, N+1)
	for i := range ok {
		ok[i] = true
	}
	for i := 1; i < N; i++ {
		if A[i-1] == A[i] {
			ok[i] = false
		}
	}
	for S := 1; S < (1 << N); S++ {
		cnt := 0
		for i := 0; i < N; i++ {
			if (S & (1 << i)) != 0 {
				cnt++
			}
		}
		if !ok[cnt] {
			continue
		}
		for T := S; T > 0; {
			T = (T - 1) & S
			if (w[S]&T) == w[S] && dp[T] {
				dp[S] = true
			}
		}
	}
	if dp[(1<<N)-1] {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
