package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	var a [1000000]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	mp := make(map[int]int)
	p := 0
	ans := 0
	for i := 0; i < n; i++ {
		mp[a[i]]++
		for p < i && k < len(mp) {
			if mp[a[p]] > 1 {
				mp[a[p]]--
			} else {
				delete(mp, a[p])
			}
			p++
		}
		ans = max(ans, i-p+1)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
