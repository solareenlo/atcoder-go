package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscan(in, &l, &n)

	x := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}

	x[n+1] = l
	pre := make([]int, n+2)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + x[i]
	}

	ans := 0
	for i, j := 1, n-1; i <= n; i, j = i+1, j-1 {
		t1 := i*l + 2*pre[(i-1)/2] + x[(i+1)/2] - 2*(pre[i]-pre[i-(i+1)/2])
		t2 := j*l + 2*(pre[i+j/2]-pre[i-1]) - x[i+(j+1)/2] - 2*(pre[n]-pre[i+(j+1)/2])
		ans = max(ans, max(t1, t2))
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
