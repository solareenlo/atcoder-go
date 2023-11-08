package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var cH, cW [2000][26]int

	var H, W int
	fmt.Fscan(in, &H, &W)
	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < W; j++ {
			cH[i][s[j]-'a']++
			cW[j][s[j]-'a']++
		}
	}
	f := true
	mH, mW := H, W
	for f {
		f = false
		dH := make([]pair, 0)
		dW := make([]pair, 0)
		for i := 0; i < H; i++ {
			for j := 0; j < 26; j++ {
				if mW > 1 && cH[i][j] == mW {
					dH = append(dH, pair{i, j})
					f = true
				}
			}
		}
		for i := 0; i < W; i++ {
			for j := 0; j < 26; j++ {
				if mH > 1 && cW[i][j] == mH {
					dW = append(dW, pair{i, j})
					f = true
				}
			}
		}
		for _, di := range dH {
			i, j := di.x, di.y
			cH[i][j] = 0
			mH--
			for k := 0; k < W; k++ {
				if cW[k][j] > 0 {
					cW[k][j]--
				}
			}
		}
		for _, di := range dW {
			i, j := di.x, di.y
			cW[i][j] = 0
			mW--
			for k := 0; k < H; k++ {
				if cH[k][j] > 0 {
					cH[k][j]--
				}
			}
		}
	}
	fmt.Println(mH * mW)
}
