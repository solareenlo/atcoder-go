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
	dp := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !dp[i] {
			fmt.Println(i)
			for j := 2 * i; j <= n; j += i {
				dp[j] = true
			}
		}
	}
}
