package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var H, W int
var S [3][1010][]string

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &H, &W)
	for i := 0; i < H; i++ {
		var t string
		fmt.Fscan(in, &t)
		S[0][i] = strings.Split(t, "")
		S[1][i] = make([]string, 1010)
		S[2][i] = make([]string, 1010)
	}
	fmt.Println(solve(0))
}

func solve(b int) int {
	ylo := H
	yhi := -1
	xlo := W
	xhi := -1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[b][i][j] == "." {
				ylo = min(ylo, i)
				yhi = max(yhi, i)
				xlo = min(xlo, j)
				xhi = max(xhi, j)
			}
		}
	}
	if yhi == -1 {
		return 0
	}
	if yhi-ylo < H/2 {
		if xhi-xlo < W/2 {
			return 1
		}
		return 2
	}
	if xhi-xlo < W/2 {
		return 2
	}
	if b >= 2 {
		return 2
	}

	ans := 4
	for w := 0; w < 4; w++ {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				S[b+1][i][j] = S[b][i][j]
			}
		}
		var fy, fx int
		if w == 0 || w == 2 {
			fy = ylo
		} else {
			fy = yhi - H/2 + 1
		}
		if w == 0 || w == 1 {
			fx = xlo
		} else {
			fx = xhi - W/2 + 1
		}

		for i := 0; i < H/2; i++ {
			for j := 0; j < W/2; j++ {
				S[b+1][i+fy][j+fx] = "#"
			}
		}

		ans = min(ans, solve(b+1)+1)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
