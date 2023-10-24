package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s [105]int

	var n, m, a, b int
	fmt.Fscan(in, &n, &m, &a, &b)
	for i := 1; i <= n; i++ {
		s[i] = b
	}
	for i := 1; i <= m; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		for j := l; j <= r; j++ {
			s[j] = a
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans += s[i]
	}
	fmt.Println(ans)
}
