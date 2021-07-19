package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Scan(&n, &k)

	c := make([]int, n+1)
	candy := make(map[int]int)
	var s, res int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
		if candy[c[i]] == 0 {
			s++
		}
		candy[c[i]]++
		if i > k {
			candy[c[i-k]]--
			if candy[c[i-k]] == 0 {
				s--
			}
		}
		if i >= k {
			res = max(res, s)
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
