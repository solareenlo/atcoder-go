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

	var T int
	fmt.Fscan(in, &T)

	a := make([]int, 2000005)
	for i := 0; i < T; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		sum := 0
		ans := 0
		for i := 1; i <= n; i++ {
			var x, y int
			fmt.Fscan(in, &x, &y)
			if i == 1 {
				ans = x
			}
			if x < 0 {
				mx := min(sum/(-x), y)
				mx = max(mx, 1)
				ans = max(ans, a[i-1]+sum*mx+x*(mx+1)*mx/2)
			}
			a[i] = a[i-1] + sum*y + x*(y+1)*y/2
			ans = max(ans, a[i])
			sum = sum + x*y
		}
		fmt.Fprintln(out, ans)
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
