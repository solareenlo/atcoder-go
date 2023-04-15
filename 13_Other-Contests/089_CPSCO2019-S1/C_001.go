package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, k, mny, cnt, ans int
var a []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &k)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	c := make([]int, 1)
	dfs(c)
	fmt.Println(ans)
}

func dfs(b []int) {
	if len(b) == k+1 {
		mny = 0
		cnt = 0
		for i := 1; i <= k; i++ {
			mny += a[b[i]-1]
		}
		for mny > 0 {
			if mny%10 > 4 {
				cnt += mny%10 - 4
			} else {
				cnt += mny % 10
			}
			mny /= 10
		}
		if ans == 0 {
			ans = cnt
		} else {
			ans = min(ans, cnt)
		}
		return
	}

	b = append(b, b[len(b)-1])
	b[len(b)-1]++
	for b[len(b)-1] <= n {
		dfs(b)
		b[len(b)-1]++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
