package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)
	A := make([][]int, h+1)
	for i := range A {
		A[i] = make([]int, w+1)
	}
	A[1][1] = 1
	for i := 1; i <= h; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 1; j <= w; j++ {
			if c[j-1] == '.' {
				if A[i-1][j] > 0 {
					A[i][j] += A[i-1][j]
				}
				if A[i][j-1] > 0 {
					A[i][j] += A[i][j-1]
				}
			}
		}
	}
	fmt.Println(A[h][w])
}
