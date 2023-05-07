package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)

	var P [8][10000]int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &P[i][j])
		}
	}

	ans := 0
	var cnt, vis [80001]int
	for i := 1; i < 1<<H; i++ {
		h := popcount(uint32(i))
		for j := 0; j < W; j++ {
			x := 0
			for k := 0; k < H; k++ {
				if ((i >> k) & 1) != 0 {
					if x == 0 || x == P[k][j] {
						x = P[k][j]
					} else {
						x = -1
					}
				}
			}
			if x > 0 {
				if vis[x] < i {
					cnt[x] = 0
				}
				vis[x] = i
				cnt[x] += h
				if ans < cnt[x] {
					ans = cnt[x]
				}
			}
		}
	}
	fmt.Println(ans)
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
