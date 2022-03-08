package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	ans := 0
	c := [30][2]int{}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		ans += a[i]
		for j := 0; j < 30; j++ {
			c[j][a[i]>>j&1]++
		}
	}

	for i := 1; i <= n; i++ {
		res := 0
		for j := 0; j < 30; j++ {
			res += c[j][1-(a[i]>>j&1)] << j
		}
		ans = max(ans, res)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
