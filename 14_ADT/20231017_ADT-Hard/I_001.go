package main

import "fmt"

func main() {
	const MOD = 998244353

	var f [200200]int

	var n int
	fmt.Scan(&n)
	var ch string
	fmt.Scan(&ch)
	ch = " " + ch + " "
	f[0] = 1
	f[1] = int(ch[1] - 48)
	sum := f[0] + f[1]
	for i := 2; i <= n; i++ {
		f[i] = f[i-1]*10%MOD + sum*int(ch[i]-48)%MOD
		f[i] %= MOD
		sum += f[i]
		sum %= MOD
	}
	fmt.Println(f[n])
}
