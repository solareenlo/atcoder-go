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

	var num [5][55]int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= n; j++ {
			var t int
			fmt.Fscan(in, &t)
			num[i][t%46]++
		}
	}

	ans := 0
	for i := 0; i < 46; i++ {
		for j := 0; j < 46; j++ {
			ans += num[1][i] * num[2][j] * num[3][(138-i-j)%46]
		}
	}
	fmt.Println(ans)
}
