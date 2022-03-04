package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	ans int
	a   = [250025]int{}
	b   = [250025]int{}
	c   = [250025]int{}
)

func dfs(L, R, bit int) {
	if L > R {
		return
	}
	if bit == -1 {
		for i := 0; i <= 17; i++ {
			cnt0 := 0
			cnt1 := 0
			for j := L; j <= R; j++ {
				if (a[j]>>i)&1 != 0 {
					cnt1++
				} else {
					cnt0++
				}
			}
			ans += cnt0 * cnt1 * (1 << i)
		}
		return
	}
	i := L
	j := R
	for i < j {
		for i <= R && (c[i]>>bit&1) == 0 {
			i++
		}
		for i < j && (c[j]>>bit)&1 != 0 {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
			b[i], b[j] = b[j], b[i]
			c[i], c[j] = c[j], c[i]
		}
	}
	for k := 0; k <= 17; k++ {
		cnt := [2][2][2][2]int{}
		for p := L; p <= R; p++ {
			tmp := 0
			if p >= i {
				tmp = 1
			}
			cnt[tmp][0][a[p]>>bit&1][a[p]>>k&1]++
			cnt[tmp][1][a[p]>>bit&1][b[p]>>k&1]++
		}
		for p := 0; p < 2; p++ {
			for q := 0; q < 2; q++ {
				ans += cnt[0][p^q][p][0] * cnt[1][p^q][q][1] * (1 << k)
				ans += cnt[0][p^q][p][1] * cnt[1][p^q][q][0] * (1 << k)
			}
		}
	}
	dfs(L, i-1, bit-1)
	dfs(i, R, bit-1)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := 1; i <= n; i++ {
		c[i] = a[i] ^ b[i]
	}
	dfs(1, n, 17)
	fmt.Println(ans)
}
