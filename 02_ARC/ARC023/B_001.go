package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var r, c, d int
	fmt.Fscan(in, &r, &c, &d)

	a := make([][]int, r)
	for i := range a {
		a[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	maxi := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i+j <= d && (i+j)%2 == d%2 {
				maxi = max(maxi, a[i][j])
			}
		}
	}

	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
