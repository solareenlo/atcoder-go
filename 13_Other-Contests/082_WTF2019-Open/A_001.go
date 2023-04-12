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
	var a [101]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans := make([]int, 0)
	var pre int
	for i := 0; i <= k; i++ {
		m := 100
		var x int
		for j := 1; j <= n; j++ {
			if j != pre && a[j] < m {
				x = j
				m = a[x]
			}
		}
		if m == 0 {
			fmt.Println(-1)
			return
		}
		for j := 1; j <= n; j++ {
			if j != pre && j != x {
				ans = append(ans, j)
				a[j]--
			}
		}
		ans = append(ans, x)
		a[x]--
		pre = x
	}
	fmt.Println(n + (n-1)*k)
	for i := range ans {
		fmt.Print(ans[i], " ")
	}
	fmt.Println()
}
