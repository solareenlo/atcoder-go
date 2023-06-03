package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	var D [4]int
	for N > 0 {
		N--
		var A, B int
		fmt.Fscan(in, &A, &B)
		D[0] += max(+A+B, 0)
		D[1] += max(+A-B, 0)
		D[2] += max(-A+B, 0)
		D[3] += max(-A-B, 0)
	}
	ans := 0
	for i := range D {
		ans = max(ans, D[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
