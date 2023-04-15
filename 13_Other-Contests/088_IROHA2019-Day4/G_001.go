package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	type pair struct {
		x, y int
	}

	var N, K int
	fmt.Fscan(in, &N, &K)
	N *= 2
	var Pair [600]int
	for i := 0; i < N; i++ {
		Pair[i] = -1
	}
	for i := 0; i < K; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		Pair[a] = b
		Pair[b] = a
	}
	n := 0
	var To [600]int
	for i := 0; i < N; i++ {
		if Pair[i] == -1 {
			To[n] = i
			n++
		}
	}
	var DP1 [600][600]pair
	var DP2 [600][600]pair
	for i := 0; i < n-1; i++ {
		c := 0
		for j := To[i] + 1; j < To[i+1]; j++ {
			if Pair[j] < To[i] || To[i+1] < Pair[j] {
				c++
			}
		}
		DP1[i][i+1] = pair{c, 1}
		DP2[i][i+1] = pair{c, 1}
	}
	for d := 3; d < N; d += 2 {
		for i := 0; i+d < n; i++ {
			count := pair{100000000, 0}
			for j := i + 1; j < i+d; j += 2 {
				if DP1[i][j].x+DP2[j+1][i+d].x < count.x {
					count = pair{DP1[i][j].x + DP2[j+1][i+d].x, DP1[i][j].y * DP2[j+1][i+d].y % mod}
				} else if DP1[i][j].x+DP2[j+1][i+d].x == count.x {
					count.y = (count.y + DP1[i][j].y*DP2[j+1][i+d].y) % mod
				}
			}
			c := 0
			for j := To[i] + 1; j < To[i+d]; j++ {
				if Pair[j] != -1 && (Pair[j] < To[i] || To[i+d] < Pair[j]) {
					c++
				}
			}
			DP1[i][i+d] = pair{c + DP2[i+1][i+d-1].x, DP2[i+1][i+d-1].y}
			if DP1[i][i+d].x == count.x {
				DP2[i][i+d] = pair{count.x, (DP1[i][i+d].y + count.y) % mod}
			} else if DP1[i][i+d].x < count.x {
				DP2[i][i+d] = DP1[i][i+d]
			} else {
				DP2[i][i+d] = count
			}
		}
	}
	fmt.Println(DP2[0][n-1].y)
}
