package main

import "fmt"

func main() {

	var S string
	fmt.Scan(&S)

	mod := 998244353
	sum, k, pow2 := 0, 0, 1
	for i := 0; i < len(S); i++ {
		k = 2*k%mod + sum
		k %= mod
		sum = sum*10%mod + pow2*int(S[i]-'0')%mod
		sum %= mod
		pow2 *= 2
		pow2 %= mod
	}
	fmt.Println((sum + k) % mod)
}
