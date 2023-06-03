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
	a := make([]int, n)
	b := make([]int, n)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 2; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	dp[0] = 0
	dp[1] = a[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+a[i], dp[i-2]+b[i])
	}

	ans := make([]int, 0)
	place := n - 1
	for {
		ans = append(ans, place)
		if place == 0 {
			break
		}
		if dp[place-1]+a[place] == dp[place] {
			place--
		} else {
			place -= 2
		}
	}
	reverseOrderInt(ans)
	fmt.Println(len(ans))
	for i := range ans {
		fmt.Printf("%d ", ans[i]+1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
