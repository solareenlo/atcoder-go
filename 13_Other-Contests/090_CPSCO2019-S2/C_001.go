package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	var s string
	fmt.Fscan(in, &n, &k, &s)
	x := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			x[i+1] = x[i] + 1
		} else {
			x[i+1] = x[i] - 1
		}
	}
	sort.Ints(x)
	x = reverseOrderInt(x)

	ans := 0
	for i := 0; i < k; i++ {
		ans += x[i]
	}
	fmt.Println(ans)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
