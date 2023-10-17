package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INFL = math.MaxInt

var h, w int
var b, c, D [500][500]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &h, &w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(in, &b[i][j])
			c[i][j] = b[i][j]
		}
	}
	res1 := solve()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			b[i][j] = c[j][i]
		}
	}
	h, w = w, h
	res2 := solve()
	fmt.Println(max(res1, res2))
}

func solve() int {
	var sum [500][500]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			sum[i][j+1] = sum[i][j] + b[i][j]
		}
	}
	for l := 0; l < w; l++ {
		for r := l; r < w; r++ {
			res := -INFL
			Min := 0
			s := 0
			for k := 0; k < h; k++ {
				s += sum[k][r+1] - sum[k][l]
				res = max(res, s-Min)
				Min = min(Min, s)
			}
			D[l][r] = res
		}
	}
	for I := 2; I <= w; I++ {
		for i := 0; i < w-I+1; i++ {
			j := i + I - 1
			D[i][j] = max(D[i][j], D[i+1][j], D[i][j-1])
		}
	}
	Max := -INFL
	for i := 0; i < w-1; i++ {
		Max = max(Max, D[0][i]+D[i+1][w-1])
	}
	return Max
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
