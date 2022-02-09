package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	ans := max(n, m)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	a := make([]int, 2005)
	b := make([]int, 2005)
	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			if i != 0 && (s[i-1][j+1]^s[i][j]^s[i-1][j]^s[i][j+1]) == 0 {
				a[j]++
			} else {
				a[j] = 1
			}
		}
		t := 0
		for j := 0; j <= m-1; j++ {
			for t > 0 && a[b[t-1]] >= a[j] {
				t--
				x := b[t]
				if t == 0 {
					ans = max(ans, (j+1)*a[x])
				} else {
					ans = max(ans, (j-b[t-1])*a[x])
				}
			}
			b[t] = j
			t++
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
