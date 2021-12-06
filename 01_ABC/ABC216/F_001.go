package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 998244353
const maxi = 5000

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Scan(&n)

	type pair struct{ x, y int }
	data := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i].x)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i].y)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].x < data[j].x
	})

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = data[i].x
		b[i] = data[i].y
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxi+1)
	}
	dp[0][0] = 1

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < maxi+1; j++ {
			dp[i+1][j] = dp[i][j]
			if b[i] <= j {
				dp[i+1][j] += dp[i][j-b[i]]
				dp[i+1][j] %= mod
			}
			if j <= a[i]-b[i] {
				res += dp[i][j]
				res %= mod
			}
		}
	}
	fmt.Println(res)
}
