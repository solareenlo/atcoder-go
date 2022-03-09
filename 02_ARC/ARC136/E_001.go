package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	c := make([]int, 1000005)
	fmt.Fscan(in, &n, &c[1])

	v := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if v[i] == 0 {
			for j := i; j <= n; j += i {
				if v[j] == 0 {
					v[j] = i
				}
			}
		}
		var x int
		fmt.Fscan(in, &x)
		if i&1 != 0 {
			c[i-v[i]+1] += x
			c[min(i+v[i], n+1)] -= x
		} else {
			c[i] += x
			c[i+1] -= x
		}
	}

	for i := 1; i <= n; i++ {
		c[i] += c[i-1]
	}

	res := 0
	for i := 1; i < n+2; i++ {
		res = max(res, c[i])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
