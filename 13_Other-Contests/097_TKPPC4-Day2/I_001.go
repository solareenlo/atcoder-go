package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var M int
	fmt.Fscan(in, &M)

	mp := make(map[int]int)
	for j := 0; j < M; j++ {
		var a int
		fmt.Fscan(in, &a)
		for i := 2; i*i <= a; i++ {
			for a%i == 0 {
				mp[i]++
				a /= i
			}
		}
		if a > 1 {
			mp[a]++
		}
	}

	keys := make([]int, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	c := make([]int, 0)
	for _, k := range keys {
		c = append(c, mp[k])
	}

	n := len(c)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		ndp := make([]int, n+1)
		for j := 0; j < n+1; j++ {
			ndp[j] = (ndp[j] + (dp[j]*j%mod)*c[i]%mod) % mod
			ndp[j] = (ndp[j] + dp[j]) % mod
			if j != n {
				ndp[j+1] = (ndp[j+1] + dp[j]*c[i]%mod) % mod
			}
		}
		dp, ndp = ndp, dp
	}

	ans := 1
	t := 1
	for i := 0; i < n; i++ {
		t *= c[i] + 1
		t %= 1000000006
	}
	ans = powMod(2, t)
	for i := 0; i < n+1; i++ {
		ans = (ans - dp[i]*2%mod + mod) % mod
	}

	fmt.Println(ans)
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
