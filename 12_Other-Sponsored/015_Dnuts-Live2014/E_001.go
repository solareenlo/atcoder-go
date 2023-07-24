package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	v := make([][]pair, 20)
	var dp [1 << 16]int

	var n, e int
	fmt.Fscan(in, &n, &e)

	for i := 0; i < e; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		if b != a {
			v[b] = append(v[b], pair{a, c})
		}
	}

	for i := 1; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if ((i >> j) & 1) != 0 {
				tmp := dp[i^(1<<j)]
				for _, p := range v[j] {
					if ((i >> p.x) & 1) != 0 {
						tmp += p.y
					}
				}
				dp[i] = max(dp[i], tmp)
			}
		}
	}
	fmt.Println(dp[(1<<n)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
