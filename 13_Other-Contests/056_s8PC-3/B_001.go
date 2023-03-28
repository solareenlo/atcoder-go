package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W, K int
	fmt.Fscan(in, &H, &W, &K)
	G := make([][]string, H)
	G2 := make([][]string, H)
	for i := 0; i < H; i++ {
		var t string
		fmt.Fscan(in, &t)
		G2[i] = strings.Split(t, "")
	}
	ans := 0
	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			for i := 0; i < H; i++ {
				G[i] = make([]string, W)
				for j := 0; j < W; j++ {
					G[i][j] = G2[i][j]
				}
			}
			for i := x - 1; i >= 0; i-- {
				G[i+1][y] = G[i][y]
			}
			G[0][y] = "-"
			scr, ss := 0, 1
			for {
				ok := false
				for i := 0; i < H; i++ {
					for j := 0; j+K <= W; {
						if G[i][j] == "-" {
							j++
							continue
						}
						k := 1
						for j+k < W && G[i][j+k] != "-" && G[i][j] == G[i][j+k] {
							k++
						}
						if k >= K {
							scr += int(G[i][j][0]-'0') * k * ss
							for l := 0; l < k; l++ {
								G[i][j+l] = "-"
							}
							ok = true
						}
						j += k
					}
				}
				for i := 0; i < W; i++ {
					for j, k := H-1, 0; j >= 0; j-- {
						if G[j][i] == "-" {
							k++
						} else {
							c := G[j][i]
							G[j][i] = "-"
							G[j+k][i] = c
						}
					}
				}
				ss *= 2
				if !ok {
					break
				}
			}
			ans = max(ans, scr)
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
