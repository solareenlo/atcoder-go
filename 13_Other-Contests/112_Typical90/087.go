package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p, k int
	fmt.Fscan(in, &n, &p, &k)

	var a [41][41]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	var binary_search func(int) int
	binary_search = func(key int) int {
		l := 1
		r := int(1e10)
		for ii := 0; ii < 40; ii++ {
			mid := (l + r) / 2
			var dp [41][41]int
			for i := 0; i < 41; i++ {
				for j := 0; j < 41; j++ {
					dp[i][j] = a[i][j]
				}
			}
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					if dp[j][k] == -1 {
						dp[j][k] = mid
					}
				}
			}
			for k := 0; k < n; k++ {
				for i := 0; i < n; i++ {
					for j := 0; j < n; j++ {
						dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
					}
				}
			}
			cnt := 0
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					if dp[i][j] <= p {
						cnt++
					}
				}
			}
			if cnt <= key {
				r = mid
			} else {
				l = mid
			}
		}
		return r
	}

	ans := binary_search(k-1) - binary_search(k)
	if ans >= 1e9 {
		fmt.Println("Infinity")
	} else {
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
