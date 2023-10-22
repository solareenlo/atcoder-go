package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	var dp [2][300000]int
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < a[i]*2; j++ {
			tmp := 0
			if i&1 == 0 {
				tmp = 1
			}
			dp[tmp][j] = (dp[i&1][j] + dp[i&1][j^a[i]]) % MOD
		}
	}
	fmt.Println(dp[n&1][k])
}
