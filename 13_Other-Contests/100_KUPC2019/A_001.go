package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, X int
	fmt.Fscan(in, &N, &X)

	var a [1005]int
	mx := 0
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &a[i])
		mx = max(mx, a[i])
	}
	cnt := 0
	for i := 1; i <= N; i++ {
		if (a[i] + X) >= mx {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
