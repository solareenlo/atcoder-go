package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	sum := 0
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		sum |= (1 << x)
	}
	dp := make([]float64, 1<<20)
	for mask := 1; mask < (1 << 16); mask++ {
		tmp := 2e18
		for i := 0; i < 16; i++ {
			val := 0.0
			cnt := 0
			for j := i - 1; j <= i+1; j++ {
				if j >= 0 && j < 16 && (mask>>j)&1 != 0 {
					cnt++
					val += dp[mask^(1<<j)]
				}
			}
			if cnt >= 0 {
				val2 := (val + 3) / float64(cnt)
				tmp = math.Min(tmp, val2)
			}
		}
		dp[mask] = tmp
	}
	fmt.Println(dp[sum])
}
