package main

import (
	"bufio"
	"fmt"
	"os"
)

var dp [100000][3]int
var st = make(map[int]bool)

func clear() {
	for c := range st {
		for i := 0; i <= 2; i++ {
			dp[c][i] = 0
		}
	}
	st = make(map[int]bool)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	ans := 0
	for i := 0; i < n; i++ {
		var s, c int
		fmt.Fscan(in, &s, &c)
		c--
		if dp[c][2] > s {
			ans++
			clear()
		} else {
			st[c] = true
			if dp[c][1] < s {
				dp[c][1] = s
			}
			if dp[c][1] > s && dp[c][2] < s {
				dp[c][2] = s
			}
		}
	}
	fmt.Println(ans)
}
