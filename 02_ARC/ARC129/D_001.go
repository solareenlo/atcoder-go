package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	sum := 0
	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	if sum != 0 {
		fmt.Println(-1)
		return
	}

	d := make([]int, N+1)
	for i := 1; i < N; i++ {
		d[i] = d[i-1] + a[i]
	}
	s := make([]int, N+1)
	for i := 1; i < N; i++ {
		s[i] = s[i-1] + d[i]
	}
	if s[N-1]%N != 0 {
		fmt.Println(-1)
		return
	}

	D := -s[N-1] / N
	mx := 0
	for i := 1; i < N; i++ {
		mx = max(mx, -(D*i + s[i]))
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans += mx
		mx += D + d[i]
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
