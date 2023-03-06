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
	S := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &S[i])
	}
	pw := make([]int, max(H, W)+1)
	pw[0] = 1
	for i := 1; i <= max(H, W); i++ {
		pw[i] = pw[i-1] * 311
	}
	ans := 0
	for i := 0; i < W; i++ {
		h := make([]int, H)
		rh := make([]int, H)
		for j := 0; j < H; j++ {
			h[j] = int(S[j][i])
			rh[j] = int(S[j][i])
		}
		for j := i + 2; j <= W; j++ {
			for k := 0; k < H; k++ {
				h[k] = h[k]*311 + int(S[k][j-1])
				rh[k] = rh[k] + int(S[k][j-1])*pw[j-i-1]
			}
			ps := make([]int, H+1)
			ps[0] = 1
			hs := make([]int, H+1)
			rhs := make([]int, H+1)
			for k := 0; k < H; k++ {
				hs[k+1] = hs[k]*pw[j-i] + h[k]
				ps[k+1] = ps[k] * pw[j-i]
			}
			for k := H - 1; k >= 0; k-- {
				rhs[k] = rhs[k+1]*pw[j-i] + rh[k]
			}
			for k := 1; k <= 2*H-3; k++ {
				l := -1
				r := (k + 1) / 2
				for r-l > 1 {
					m := (l + r) >> 1
					p := k - m + 1
					if p > H {
						l = m
					} else {
						ha := hs[p] - hs[m]*ps[p-m]
						hb := rhs[m] - rhs[p]*ps[p-m]
						if ha == hb {
							r = m
						} else {
							l = m
						}
					}
				}
				ans += (k+1)/2 - r
			}
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
