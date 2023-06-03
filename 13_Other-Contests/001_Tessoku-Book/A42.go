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

	var n, k int
	fmt.Fscan(in, &n, &k)
	A := make([]pair, n)
	for i := range A {
		fmt.Fscan(in, &A[i].x, &A[i].y)
	}
	ans := 1
	for i := 1; i <= 100-k; i++ {
		for j := 1; j <= 100-k; j++ {
			c := 0
			for _, tmp := range A {
				a := tmp.x
				b := tmp.y
				if a >= i && a <= i+k && b >= j && b <= j+k {
					c++
				}
			}
			ans = max(c, ans)
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
