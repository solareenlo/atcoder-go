package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l, k int
	fmt.Fscan(in, &n, &l, &k)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	ans := 0
	for b := 0; b < 1<<l; b++ {
		cnt := 0
		for i := 0; i < l; i++ {
			if (b & (1 << i)) != 0 {
				cnt++
			}
		}
		if cnt == k {
			mp := make(map[string]int)
			for i := 0; i < n; i++ {
				var t string
				for j := 0; j < l; j++ {
					if (b & (1 << j)) == 0 {
						t += string(s[i][j])
					}
				}
				mp[t]++
			}
			for _, x := range mp {
				ans = max(ans, x)
			}
		}
		cnt++
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
