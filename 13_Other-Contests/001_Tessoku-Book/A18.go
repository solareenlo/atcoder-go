package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [10010]bool

	var n, s int
	fmt.Fscan(in, &n, &s)
	dp[0] = true
	for i := 1; i <= n; i++ {
		var t int
		fmt.Fscan(in, &t)
		for j := s; j >= t; j-- {
			if dp[j-t] {
				dp[j] = true
			}
		}
	}
	if dp[s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
