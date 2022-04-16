package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)

	A := m + 1
	B := n + 1
	ans := 0
	for i, j := 1, 1; i <= n || j <= m; {
		if i <= n && (j > m || a[i] < b[j]) {
			ans += a[i] * A
			i++
			B--
		} else {
			ans += b[j] * B
			j++
			A--
		}
	}
	fmt.Println(ans)
}
