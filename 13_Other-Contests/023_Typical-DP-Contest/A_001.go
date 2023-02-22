package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	dp := make([]bool, 11000)
	dp[0] = true

	for i := 1; i <= N; i++ {
		var P int
		fmt.Fscan(in, &P)
		for j := i * 100; j >= 0; j-- {
			if !dp[j] {
				continue
			}
			dp[j+P] = true
		}
	}

	ans := 0
	for i := 0; i <= N*100; i++ {
		if dp[i] {
			ans++
		}
	}

	fmt.Println(ans)
}
