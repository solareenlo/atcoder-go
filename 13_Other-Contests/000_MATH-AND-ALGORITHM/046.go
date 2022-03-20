package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var r, c, sy, sx, gy, gx int
	fmt.Fscan(in, &r, &c, &sy, &sx, &gy, &gx)

	leng := [51][51]int{}
	for i := range leng {
		for j := range leng[i] {
			leng[i][j] = 100000
		}
	}

	S := make([]string, r)
	for i := range S {
		fmt.Fscan(in, &S[i])
	}

	X := make([]int, 0)
	Y := make([]int, 0)
	X = append(X, sx-1)
	Y = append(Y, sy-1)
	leng[sy-1][sx-1] = 0
	dy := []int{1, -1, 0, 0}
	dx := []int{0, 0, 1, -1}
	for len(X) > 0 {
		x := X[0]
		y := Y[0]
		X = X[1:]
		Y = Y[1:]
		for i := 0; i < 4; i++ {
			if S[y+dy[i]][x+dx[i]] == '#' {
				continue
			}
			if leng[y+dy[i]][x+dx[i]] > 1+leng[y][x] {
				leng[y+dy[i]][x+dx[i]] = 1 + leng[y][x]
				Y = append(Y, y+dy[i])
				X = append(X, x+dx[i])
			}
		}
	}
	fmt.Println(leng[gy-1][gx-1])
}
