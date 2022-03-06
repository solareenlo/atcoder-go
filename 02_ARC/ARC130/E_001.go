package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, K int
	fmt.Fscan(in, &n, &K)

	a := make([]int, K+1)
	for i := 1; i <= K; i++ {
		fmt.Fscan(in, &a[i])
	}

	pre := 0
	N := 0
	lst := make([]int, 300300)
	len := make([]int, 300300)
	for i := 1; i <= K; i++ {
		x := a[i]
		pre = max(pre, lst[x])
		if lst[x] == 0 {
			N++
		}
		lst[x] = i
		if i-pre == N {
			len[i] = len[pre] + 1
		} else {
			len[i] = -1e9
		}
	}

	mn := 1 << 60
	pos := -1
	for i := pre; i <= K; i++ {
		if len[i] >= 0 && len[i] <= mn {
			mn = len[i]
			pos = i
		}
	}

	if pos < 0 {
		fmt.Fprintln(out, -1)
		return
	}

	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = mn + 1
	}
	for i := 1; i <= pos; i++ {
		ans[a[i]]--
	}
	for i := 1; i <= n; i++ {
		fmt.Print(ans[i], " ")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
