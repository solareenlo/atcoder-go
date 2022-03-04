package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	k int
	a = [200200]int{}
)

func f(x int) int {
	return max(0, (a[x]-1)*k+1)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n, &k)

	m := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		m += a[i]
	}

	mx := 0
	b := make([]int, m+1)
	for i := 1; i <= n; i++ {
		tmp := f(i)
		if tmp > m {
			fmt.Fprintln(out, -1)
			return
		}
		b[tmp]++
		mx = max(mx, tmp)
	}

	if mx+b[mx]-1 > m {
		fmt.Fprintln(out, -1)
		return
	}

	vis := make([]bool, n+1)
	ans := make([]int, m+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if !vis[j] && a[j] != 0 {
				tmp := mx
				b[f(j)]--
				a[j]--
				b[f(j)]++
				for b[mx] == 0 {
					mx--
				}
				if mx == 0 || i+mx+b[mx]-1 <= m {
					vis[j] = true
					if i >= k {
						vis[ans[i-k+1]] = false
					}
					ans[i] = j
					break
				}
				b[f(j)]--
				a[j]++
				b[f(j)]++
				mx = tmp
			}
		}
		fmt.Fprint(out, ans[i], " ")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
