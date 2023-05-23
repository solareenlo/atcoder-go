package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)
	var dp [10000]int
	for i := range dp {
		dp[i] = -n - 1
	}
	dp[x] = 0
	P := make([]int, 4)
	for k := 0; k < n; k++ {
		var a int
		fmt.Fscan(in, &a)
		for i := 0; i < 10000; i++ {
			m := 0
			for d := i; d > 0; m++ {
				P[m] = d % 10
				d /= 10
			}
			tmp := P[:m]
			sort.Ints(tmp)
			for nextPermutation(sort.IntSlice(P[:m])) {
				s := 0
				for j := 0; j < m; j++ {
					s = 10*s + P[j]
				}
				dp[s] = max(dp[s], dp[i])
			}
		}
		for i := a; i < 10000; i++ {
			dp[i-a] = max(dp[i-a], dp[i]+1)
		}
	}
	ans := 0
	for i := 0; i < 10000; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
