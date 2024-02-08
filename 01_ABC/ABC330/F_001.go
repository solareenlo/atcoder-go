package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n, k int
var x, y []int

func check(Len int) bool {
	ans := 0
	for i := 1; i <= (n >> 1); i++ {
		ans += max(0, x[n-i+1]-x[i]-Len) + max(0, y[n-i+1]-y[i]-Len)
	}
	return ans <= k
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &k)
	x = make([]int, n+1)
	y = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	sort.Ints(x[1:])
	sort.Ints(y[1:])
	l := 0
	r := int(1e9)
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
