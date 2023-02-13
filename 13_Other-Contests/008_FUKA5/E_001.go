package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for {
		var n int
		fmt.Fscan(in, &n)
		if n == 0 {
			break
		}
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		nxt := make([]int, n)
		for x := n - 1; x >= 0; x-- {
			if a[x] == -1 {
				nxt[x] = 0
			} else if a[x] == 0 {
				nxt[x] = x
			} else {
				nxt[x] = nxt[min(x+a[x], n-1)]
			}
		}
		c1 := make([]float64, n)
		c2 := make([]float64, n)
		for x := n - 2; x >= 0; x-- {
			c1[x] = 1
			for d := 1; d <= 6; d++ {
				x2 := nxt[min(x+d, n-1)]
				if x2 == 0 {
					c2[x] += 1.0 / 6
				} else {
					c1[x] += c1[x2] / 6
					c2[x] += c2[x2] / 6
				}
			}
		}
		fmt.Printf("%.9f\n", c1[0]/(1-c2[0]))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
