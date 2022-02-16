package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	rand.Seed(time.Now().UnixNano())
	ok := [13][13]bool{}
	for l := 0; l < 50000; l++ {
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		for i := 0; i < 13; i++ {
			for j := 0; j < 13; j++ {
				ok[i][j] = true
			}
		}
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				for k := 0; k < n && ok[i][j]; k++ {
					if a[k][i] != a[n-1-k][j] {
						ok[i][j] = false
					}
				}
			}
		}
		dp := make([]int, 4100)
		dp[0] = 1
		for i := 0; i < (1 << m); i++ {
			if dp[i] != 0 {
				for j := 0; j < m; j++ {
					if (i & (1 << j)) == 0 {
						for k := j + 1; k < m; k++ {
							if ok[j][k] && (i&(1<<k)) == 0 {
								dp[i|(1<<j)|(1<<k)] = 1
							}
						}
					}
				}
			}
		}
		flag := 0
		if m%2 == 0 {
			flag = dp[(1<<m)-1]
		} else {
			for i := 0; i < m; i++ {
				flag |= dp[(1<<m)-1-(1<<i)]
			}
		}
		if flag != 0 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
	return
}
