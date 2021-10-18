package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 2004

var (
	L [MAX][MAX]int
	R [MAX][MAX]int
	D [MAX][MAX]int
	U [MAX][MAX]int
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	s := make([]string, h+1)
	tmp := "#"
	for i := 0; i < w+1; i++ {
		tmp += "#"
	}
	s[0] = tmp
	for i := 1; i <= h; i++ {
		fmt.Fscan(in, &s[i])
		s[i] = "#" + s[i] + "#"
	}
	s = append(s, tmp)

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if s[i][j] != '#' {
				L[i][j] = L[i][j-1] + 1
				D[i][j] = D[i-1][j] + 1
			}
		}
	}
	for i := 1; i <= h; i++ {
		for j := w; j >= 1; j-- {
			if s[i][j] != '#' {
				R[i][j] = R[i][j+1] + 1
			}
		}
	}
	for j := 1; j <= w; j++ {
		for i := h; i >= 1; i-- {
			if s[i][j] != '#' {
				U[i][j] = U[i+1][j] + 1
			}
		}
	}

	res := 0
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			res = max(res, L[i][j]+R[i][j]+D[i][j]+U[i][j]-3)
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
