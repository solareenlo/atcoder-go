package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	lim := int(math.Sqrt(float64(M)))
	prime := make([]bool, lim+1)
	for i := range prime {
		prime[i] = true
	}
	prime[0] = false
	prime[1] = false
	for i := 2; i <= lim; i++ {
		if prime[i] {
			for j := i * 2; j <= lim; j += i {
				prime[j] = false
			}
		}
	}
	mobius := make([]int, lim+1)
	for i := range mobius {
		mobius[i] = 1
	}
	for i := 2; i <= lim; i++ {
		if prime[i] {
			for j := i; j <= lim; j += i {
				if j%(i*i) == 0 {
					mobius[j] = 0
				} else {
					mobius[j] *= -1
				}
			}
		}
	}
	allin := make([]int, lim+1)
	for i := 1; i <= lim; i++ {
		z := M / (i * i)
		sublim := lim / i
		for j := 1; j <= sublim; j++ {
			tmp := int(math.Sqrt(float64(z) / float64(j)))
			allin[tmp] = (allin[tmp] + mobius[i]) % MOD
		}
		for j := 1; j <= sublim; j++ {
			num := max(z/j+1, sublim+1) - max(z/(j+1)+1, sublim+1)
			tmp := int(math.Sqrt(float64(j)))
			allin[tmp] = (allin[tmp] + num*mobius[i]%MOD) % MOD
		}
	}
	seqs := make([]int, N+1)
	for i := 1; i <= lim; i++ {
		if allin[i] != 0 {
			mul := 1
			for j := 1; j <= N; j++ {
				mul = mul * i % MOD
				seqs[j] = (seqs[j] + (allin[i] * mul % MOD)) % MOD
			}
		}
	}

	pack := int(math.Sqrt(float64(N)))
	dp1 := make([][]int, N+1)
	dp1[0] = []int{1}
	for i := 1; i <= N; i++ {
		dp1[i] = make([]int, i+1)
		for j := 1; j <= i; j++ {
			for k := 1; k <= i-j+1 && k < pack; k++ {
				dp1[i][j] = (dp1[i][j] + dp1[i-k][j-1]*seqs[k]%MOD) % MOD
			}
		}
	}

	dp2 := make([][]int, N+1)
	for i := range dp2 {
		dp2[i] = make([]int, N/pack+1)
	}
	dp2[0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= N/pack; j++ {
			for k := pack; k <= i; k++ {
				dp2[i][j] = (dp2[i][j] + dp2[i-k][j-1]*seqs[k]%MOD) % MOD
			}
		}
	}

	fact := make([]int, N+1)
	factinv := make([]int, N+1)
	fact[0] = 1
	for i := 1; i <= N; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
	for i := 0; i <= N; i++ {
		factinv[i] = invMod(fact[i])
	}
	for i := 0; i <= N; i++ {
		for j := 0; j <= N/pack; j++ {
			dp2[i][j] = dp2[i][j] * factinv[j] % MOD
		}
	}
	total := make([]int, N+1)
	for i := 0; i <= N; i++ {
		for j := 0; j <= i; j++ {
			v := dp1[i][j] * factinv[j] % MOD
			for k := 0; k <= N/pack && j+k <= N; k++ {
				total[j+k] = (total[j+k] + v*dp2[N-i][k]%MOD) % MOD
			}
		}
	}
	ans := make([]int, N)
	for i := 1; i <= N; i++ {
		ans[N-i] = total[i] * fact[i] % MOD
	}
	for i := N - 1; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			ans[i] = (ans[i] - ((ans[j]*fact[j]%MOD)*factinv[i]%MOD)*factinv[j-i]%MOD + MOD) % MOD
		}
	}
	for i := 0; i < N; i++ {
		fmt.Println(ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func invMod(a int) int {
	return powMod(a, MOD-2)
}
