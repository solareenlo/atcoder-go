package main

import "fmt"

func main() {
	const mod = 1000000007

	var d, x int
	fmt.Scan(&d, &x)
	if x == 1 {
		fmt.Println(0)
		return
	}
	var eq, neq int
	var a, b [3]int
	a[0] = x - 1
	b[0] = x - 2
	u := (a[0]*a[0] + b[0]*b[0]%mod*(x-1)) % mod
	v := (a[0]*b[0]*2 + (x-2)*b[0]%mod*b[0]) % mod
	w := ((x-1)*2*a[0]%mod*b[0] + (x-1)*(x-2)%mod*b[0]%mod*b[0]) % mod
	z := (a[0]*a[0] + (x-2)*2*a[0]%mod*b[0] + (x*x-x*3+3)%mod*b[0]%mod*b[0]) % mod

	for i := 1; i <= d; i++ {
		j := (i + 2) % 3
		tmp := b[j] * (x - 2) % mod
		a[i%3] = (a[j]*(x-1) + (x-1)*tmp) % mod
		b[i%3] = (a[j]*(x-2) + b[j]*(x-1) + (x-2)*tmp) % mod
		if i == 1 {
			eq = x * a[1] % mod
			neq = 0
		} else {
			k := (i + 1) % 3
			new_eq := (eq*u + neq*v) % mod * a[k]
			new_neq := (eq*w + neq*z) % mod * b[k]
			eq = new_eq % mod
			neq = new_neq % mod
		}
	}
	fmt.Println((eq + neq) % mod)
}
