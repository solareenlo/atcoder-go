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

	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}

	mp := make(map[int]int)
	for i := 0; i < k; i++ {
		mp[c[i]]++
	}

	res := len(mp)
	for i := k; i < n; i++ {
		mp[c[i]]++
		mp[c[i-k]]--
		if mp[c[i-k]] == 0 {
			delete(mp, c[i-k])
		}
		res = max(res, len(mp))
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
