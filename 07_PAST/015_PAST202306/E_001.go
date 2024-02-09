package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 8

	var n, k int
	fmt.Fscan(in, &n, &k)
	var as [N]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &as[i])
	}

	var fs [N + 1]int
	fs[0] = 1
	for i := 1; i <= N; i++ {
		fs[i] = fs[i-1] * i
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += as[i]
	}

	ans := sum * (fs[n-1] / (fs[k-1] * fs[n-k]))
	fmt.Println(ans)
}
