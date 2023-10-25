package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [100100]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 0; i < k; i++ {
		a[i+1] = -1e18
	}
	for j := 0; j < n; j++ {
		var t, z int
		fmt.Fscan(in, &t, &z)
		if t != 0 {
			for i := 0; i < k; i++ {
				a[k-i] = max(a[k-i], a[k-i-1]+z)
			}
		} else {
			for i := 0; i < k; i++ {
				a[i] = max(a[i], a[i+1]+z)
			}
		}
	}
	ans := 0
	for i := 0; i < k+1; i++ {
		ans = max(a[i], ans)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
