package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	s := make([]int, n+2)
	for i := n; i >= 0; i-- {
		s[i] = s[i+1] + a[i]
	}

	ans := 0
	no := 1
	for i := 0; i <= n; i++ {
		ans += no
		no = no - a[i]
		if no < 0 {
			fmt.Println(-1)
			return
		}
		no = min(no*2, s[i+1])
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
