package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	var dp [1 << 18]int
	dp[0] = 1
	var a [18]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		for j := 0; j < 1<<i; j++ {
			dp[1<<i|j] += dp[j]
			for I := 0; I < i; I++ {
				if (j & (1 << I)) != 0 {
					continue
				}
				if a[i]+a[I] <= 2*k {
					dp[1<<i|1<<I|j] += dp[j]
				}
				for J := I + 1; J < i; J++ {
					if (j & (1 << J)) != 0 {
						continue
					}
					if a[i]+a[I]+a[J] <= 3*k {
						dp[1<<i|1<<I|1<<J|j] += dp[j]
					}
				}
			}
		}
	}
	fmt.Println(dp[(1<<n)-1])
}
