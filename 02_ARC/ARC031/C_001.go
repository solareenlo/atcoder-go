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

	L := make([]int, n+1)
	ans := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		s := 0
		for j := a; j > 0; j -= j & -j {
			s += L[j]
		}
		l := i - s
		r := n - i - a + s
		ans += min(l, r)
		for j := a; j <= n; j += j & -j {
			L[j]++
		}
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
