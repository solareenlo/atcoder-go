package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var w, h, b, a int
	fmt.Fscan(in, &w, &h, &b, &a)
	var c [1001][1001]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var x int
			fmt.Fscan(in, &x)
			if x == -1 {
				x = int(1e8)
			}
			c[i+1][j+1] = x + c[i+1][j] + c[i][j+1] - c[i][j]
		}
	}

	ans := int(1e8)
	for i := 0; i < h-a+1; i++ {
		for j := 0; j < w-b+1; j++ {
			ans = min(ans, c[i+a][j+b]-c[i+a][j]-c[i][j+b]+c[i][j])
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
