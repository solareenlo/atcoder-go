package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Ints(tmp)

	ans := a[n] - a[1]
	for i := 2; i <= n/2; i++ {
		ans += a[n-i+2] - a[i] + a[n-i+1] - a[i-1]
	}
	if n&1 != 0 {
		ans += max(a[n/2+1]-a[n/2], a[n/2+2]-a[n/2+1])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
