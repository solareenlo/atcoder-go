package main

import "fmt"

func main() {
	const MOD = 1000000007

	var pow2 [45010]int
	var f [160][160][160]int
	var C [160][160]int
	var n, d int
	fmt.Scan(&n, &d)
	pow2[0] = 1
	for i := 1; i <= n*n; i++ {
		pow2[i] = (pow2[i-1] << 1) % MOD
	}
	for i := 0; i < n+1; i++ {
		C[i][0] = 1
		C[i][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}
	f[0][1][1] = 1
	for i := 0; i < d; i++ {
		for j := i; j <= n; j++ {
			for k := 1; k <= j-i+1; k++ {
				if f[i][j][k] != 0 {
					for tmp, l := f[i][j][k], 1; l <= n-j-d+i+1; l++ {
						tmp = tmp * (pow2[k] - 1) % MOD
						tmp1 := 0
						if i+1 == d {
							tmp1 = 1
						}
						f[i+1][j+l][l] = (f[i+1][j+l][l] + (tmp*pow2[l*(l-1)/2]%MOD)*C[n-j-1][l-tmp1]%MOD) % MOD
					}
				}
			}
		}
	}
	ans := 0
	for i := d; i <= n; i++ {
		for k := 1; k <= n; k++ {
			ans = (ans + f[d][i][k]*pow2[(n-i)*(n-i-1)/2+(n-i)*k]%MOD) % MOD
		}
	}
	fmt.Println(ans)
}
