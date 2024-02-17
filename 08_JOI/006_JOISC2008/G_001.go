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

	var N, H, W int
	fmt.Fscan(in, &N, &H, &W)
	m := make(map[pair]int)
	for i := 0; i < N; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &c, &b, &d)
		for j := a; j <= b; j++ {
			for k := c; k <= d; k++ {
				m[pair{j, k}]++
			}
		}
	}
	ma := 0
	for _, x := range m {
		ma = max(ma, x)
	}
	fmt.Println(ma)
	ans := 0
	for _, x := range m {
		if x == ma {
			ans++
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
