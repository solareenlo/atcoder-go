package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 3005

	var n int
	fmt.Fscan(in, &n)
	var a, bj, cnt, inv, jc, dp [N]int
	gs := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] == -1 {
			gs++
		} else {
			bj[a[i]] = 1
			for j := a[i] + 1; j <= n; j++ {
				cnt[j]++
			}
		}
	}
	sum0 := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if a[i] != -1 && a[j] != -1 && a[i] > a[j] {
				sum0++
			}
		}
	}
	inv[0] = 1
	jc[0] = 1
	for i := 1; i <= n; i++ {
		jc[i] = jc[i-1] * i % MOD
	}
	inv[n] = powMod(jc[n], MOD-2)
	ny := powMod(gs, MOD-2)
	ny1 := powMod(gs*(gs-1), MOD-2)
	for i := n - 1; i >= 1; i-- {
		inv[i] = inv[i+1] * (i + 1) % MOD
	}
	sum1 := 0
	ans11, ans12, dcnt := 0, 0, 0
	for i := 1; i <= n; i++ {
		if a[i] == -1 {
			pr := sum1
			dsum := 0
			for j := 1; j <= n; j++ {
				if bj[j] == 0 {
					ans11 = (ans11 + cnt[j]*cnt[j]*ny) % MOD
					ans11 = (ans11 + (pr-dp[j])*cnt[j]%MOD*2*ny1) % MOD
					ans12 = (ans12 + ((gs-dsum-1)*dcnt+(gs-dcnt-1)*dsum)%MOD*cnt[j]%MOD*ny1) % MOD
					sum1 += cnt[j]
					dp[j] += cnt[j]
					dsum++
				}
			}
			dcnt++
		} else {
			for j := 1; j <= a[i]-1; j++ {
				cnt[j]++
			}
			for j := a[i] + 1; j <= n; j++ {
				cnt[j]--
			}
		}
	}
	sum1 = sum1 * ny % MOD
	sum2 := gs * (gs - 1) / 2 * inv[2] % MOD
	ans12 = (ans12 + sum1*(gs-1)*(gs-2)/2%MOD*inv[2]) % MOD
	ans22 := gs * (gs - 1) / 2 * inv[2] % MOD
	ans22 = (ans22 + gs*(gs-1)*(gs-2)%MOD*inv[3]%MOD*inv[3]%MOD*10) % MOD
	ans22 = (ans22 + gs*(gs-1)/2%MOD*(gs-2)*(gs-3)/2%MOD*inv[2]%MOD*inv[2]) % MOD
	ans00 := sum0 * sum0 % MOD
	ans01 := sum0 * sum1 % MOD
	ans02 := sum0 * sum2 % MOD
	fmt.Println((ans00 + ans11 + ans22 + ans01*2 + ans02*2 + ans12*2) * jc[gs] % MOD)
}

const MOD = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
