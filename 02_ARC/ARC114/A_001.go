package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n     int
	x         = [55]int{}
	res   int = 1 << 60
	prime     = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
)

func dfs(i, j int) {
	if i == 15 {
		for i := 0; i < n; i++ {
			if gcd(j, x[i]) == 1 {
				return
			}
		}
		res = min(res, j)
		return
	}
	dfs(i+1, j)
	dfs(i+1, j*prime[i])
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
	}

	dfs(0, 1)

	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
