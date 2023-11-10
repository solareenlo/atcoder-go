package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	const N = 200005

	var a, b, f, g, cnt [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		f[i] = INF
	}
	for x := 1; x < 1e10; x *= 10 {
		for i := 0; i < 11; i++ {
			cnt[i] = 0
		}
		for i := 0; i < n+1; i++ {
			g[i] = f[i]
			f[i] = INF
			if i != 0 {
				cnt[a[i]/x%10]++
			}
		}
		for i := 0; i < n+1; i++ {
			if i != 0 {
				cnt[a[i]/x%10]--
				cnt[a[i]/x%10+1]++
			}
			if g[i] < INF {
				for j := 0; j < 10; j++ {
					tmp := 0
					sum := 0
					for k := 0; k < 11; k++ {
						if j+k >= 10 {
							tmp += cnt[k]
							sum += (j + k - 10) * cnt[k]
						} else {
							sum += (j + k) * cnt[k]
						}
					}
					f[tmp] = min(f[tmp], g[i]+sum)
				}
			}
		}
		for i := 0; i < 11; i++ {
			cnt[i] = 0
		}
		for i := 1; i <= n; i++ {
			cnt[a[i]/x%10]++
		}
		for i := 1; i <= 9; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := 1; i <= n; i++ {
			b[cnt[a[i]/x%10]] = a[i]
			cnt[a[i]/x%10]--
		}
		for i := 1; i <= n; i++ {
			a[i] = b[n-i+1]
		}
	}
	fmt.Println(f[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
