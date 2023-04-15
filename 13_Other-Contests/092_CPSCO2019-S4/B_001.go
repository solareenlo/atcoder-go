package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)
	var s [11][11]bool
	for i := 1; i <= d; i++ {
		var h string
		fmt.Fscan(in, &h)
		for j := 0; j < n; j++ {
			if h[j] == 'o' {
				s[i][j+1] = true
			}
		}
	}
	ans := 0
	for i := 1; i < d; i++ {
		for j := i + 1; j <= d; j++ {
			cnt := 0
			for k := 1; k <= n; k++ {
				if s[i][k] || s[j][k] {
					cnt++
				}
			}
			ans = max(cnt, ans)
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
