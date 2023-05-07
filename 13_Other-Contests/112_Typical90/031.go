package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const M = 1505

	var n int
	fmt.Fscan(in, &n)

	var w, b [100005]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	var dp [55][1505]int
	for w := 0; w < 51; w++ {
		for b := 0; b < 1326; b++ {
			x := make([]bool, M)
			if w != 0 {
				x[dp[w-1][b+w]] = true
			}
			for k := 1; k <= b/2; k++ {
				x[dp[w][b-k]] = true
			}
			dp[w][b] = find(x, false)
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans ^= dp[w[i]][b[i]]
	}
	if ans != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}

func find(x []bool, y bool) int {
	for i := range x {
		if x[i] == y {
			return i
		}
	}
	return 0
}
