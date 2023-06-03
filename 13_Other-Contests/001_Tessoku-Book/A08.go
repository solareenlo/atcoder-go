package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)

	sum := make([][]int, H+1)
	for i := range sum {
		sum[i] = make([]int, W+1)
	}

	for h := 1; h <= H; h++ {
		for w := 1; w <= W; w++ {
			var X int
			fmt.Fscan(in, &X)
			sum[h][w] = sum[h][w-1] + X
		}
	}

	for h := 1; h <= H; h++ {
		for w := 1; w <= W; w++ {
			sum[h][w] += sum[h-1][w]
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	ans := make([]int, 0)
	for Q > 0 {
		Q--
		var A, B, C, D int
		fmt.Fscan(in, &A, &B, &C, &D)
		ans = append(ans, sum[C][D]-sum[A-1][D]-sum[C][B-1]+sum[A-1][B-1])
	}

	for i := range ans {
		fmt.Println(ans[i])
	}
}
