package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100001

	var as [N]int

	var a int
	fmt.Fscan(in, &a)
	for i := 1; i < N; i++ {
		as[i] = 1e18
	}
	e := 0
	for i := 1; i <= a; i++ {
		var b, c, d int
		fmt.Fscan(in, &b, &c, &d)
		e += d
		for j := N - 1; j >= 0; j-- {
			as[j] = min(as[j], as[max(0, j-d)]+max(0, c-(b+c)/2))
		}
	}
	fmt.Println(as[e/2+1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
