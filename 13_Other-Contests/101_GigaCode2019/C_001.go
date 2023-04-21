package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var d int
	fmt.Fscan(in, &d)

	var a [200000]int
	for i := 0; i < d; i++ {
		fmt.Fscan(in, &a[i])
	}

	sum := 0
	ans := int(1e18)
	for i := 0; i < d; i++ {
		var b int
		fmt.Fscan(in, &b)
		sum += a[i]
		if sum >= b {
			ans = min(ans, b)
		}
	}
	if ans == int(1e18) {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
