package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 5050

	var n int
	fmt.Fscan(in, &n)
	var a [N]float64
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var dp [N]float64
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			dp[j] = math.Max(dp[j], dp[j-1]*0.9+a[i])
		}
	}
	ans := -1e18
	o := 1.0
	cnt := 0.9
	for i := 1; i <= n; i++ {
		ans = math.Max(ans, dp[i]/o-1200.0/math.Sqrt(float64(i)))
		o += cnt
		cnt *= 0.9
	}
	fmt.Println(ans)
}
