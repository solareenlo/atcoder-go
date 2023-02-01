package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2020

var f [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	dx := []int{-1, -1, 0, 0, 1, 1}
	dy := []int{-1, 0, -1, 1, 0, 1}

	var n int
	fmt.Fscan(in, &n)
	ans := n
	var x, y [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		f[i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 0; k < 6; k++ {
				if x[i]+dx[k] == x[j] && y[i]+dy[k] == y[j] && find(i) != find(j) {
					f[find(i)] = find(j)
					ans--
				}
			}
		}
	}
	fmt.Println(ans)
}

func find(x int) int {
	if x == f[x] {
		return x
	}
	f[x] = find(f[x])
	return f[x]
}
