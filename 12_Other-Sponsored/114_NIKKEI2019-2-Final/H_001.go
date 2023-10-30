package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const N = 6000060

	var p, a, pre, lst, suf, ind, cnt, sum [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	p[0] = 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		pre[i] = lst[a[i]]
		lst[a[i]] = i
	}
	for i := 1; i <= m; i++ {
		lst[i] = n + 1
		p[i] = p[i-1] * m % MOD
	}
	for i := n; i >= 1; i-- {
		suf[i] = lst[a[i]]
		lst[a[i]] = i
	}
	ans, Max, MAX := 0, 0, 0
	for i := 1; i <= 2*n-1; i++ {
		l, r := 0, 0
		if i < Max {
			tmp := 2*MAX - i
			w := (i - tmp) / 2
			l = (tmp-ind[tmp])/2 + 1
			r = (tmp + ind[tmp]) / 2
			cnt[i] = cnt[tmp]
			sum[i] = sum[tmp]
			for r-l+1 > Max-i {
				sum[i] = (sum[i] - p[m-cnt[i]] + MOD) % MOD
				if suf[l] > r {
					cnt[i]--
				}
				l++
				if pre[r] < l {
					cnt[i]--
				}
				r--
			}
			l += w
			r += w
		} else {
			if (i & 1) != 0 {
				l = (i + 1) / 2
				r = (i + 1) / 2
				cnt[i] = 1
			} else {
				l = i / 2
				r = i/2 + 1
				if a[l] != a[r] {
					cnt[i] = 1 + 1
				} else {
					cnt[i] = 1
				}
			}
			sum[i] = p[m-cnt[i]]
		}
		for l > 1 && r < n && (suf[l-1] > r || a[i+1-suf[l-1]] == a[r+1]) && (pre[r+1] < l || a[i+1-pre[r+1]] == a[l-1]) {
			l--
			if suf[l] > r {
				cnt[i]++
			}
			r++
			if pre[r] < l {
				cnt[i]++
			}
			sum[i] = (sum[i] + p[m-cnt[i]]) % MOD
		}
		ind[i] = r - l + 1
		if i+ind[i] >= Max {
			Max = i + ind[i]
			MAX = i
		}
		ans = (ans + sum[i]) % MOD
	}
	fmt.Println(ans)
}
