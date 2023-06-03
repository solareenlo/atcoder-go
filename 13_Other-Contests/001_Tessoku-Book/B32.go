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
	a := make([]int, k)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	dp := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		for j := 0; j < k; j++ {
			if i >= a[j] && !dp[i-a[j]] {
				dp[i] = true
			}
		}
	}
	if dp[n] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
