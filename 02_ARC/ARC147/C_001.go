package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const maxn = 300010

	var n int
	fmt.Fscan(in, &n)
	l := make([]int, maxn)
	r := make([]int, maxn)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &l[i], &r[i])
	}
	tmp1 := l[1 : n+1]
	sort.Ints(tmp1)
	tmp2 := r[1 : n+1]
	sort.Ints(tmp2)
	tmp3 := l[1 : n+1]
	reverseOrder(tmp3)

	ans := 0
	for i := 1; i <= n/2; i++ {
		ans = ans + max(0, l[i]-r[i])*(n-2*(i-1)-1)
	}
	fmt.Println(ans)
}

func reverseOrder(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
