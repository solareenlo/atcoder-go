package main

import "fmt"

const mod = 998244353
const Nitwo = 499122177

func Find(x, t int) int {
	for i := 0; i < t; i++ {
		x = x + (x+1)/2
	}
	return x
}

func main() {
	var n, K int
	fmt.Scan(&n, &K)

	m := n
	for i := 1; i <= K; i++ {
		m = m - (m+2)/3
	}
	if m == 0 {
		fmt.Println(0)
		return
	}

	if K <= 40 {
		X := K / 2
		Y := (K + 1) / 2
		maxn1 := (1 << X)
		maxn2 := (1 << Y)

		d := 1
		d2 := 1
		for i := 1; i <= X; i++ {
			d *= 3
		}
		for i := 1; i <= Y; i++ {
			d2 *= 3
		}

		t := m / maxn1
		Sum := 0
		for i := 0; i < maxn2; i++ {
			Sum = (Sum + Find(i*d, Y)) % mod
		}
		Sum2 := 0
		for i := 0; i < t%maxn2; i++ {
			Sum2 = (Sum2 + Find(i*d, Y)) % mod
		}

		t4 := (t / maxn2) % mod * ((t/maxn2 - 1) % mod) % mod * Nitwo % mod * (d % mod) % mod * (d2 % mod) % mod

		h := make([]int, 2000200)
		for i := 0; i < maxn2; i++ {
			h[i*d%maxn2] = Sum2
			Sum2 = (Sum2 - Find(i*d%maxn2, Y)%mod + Find(i*d%maxn2+t%maxn2*d, Y) + mod) % mod
			Sum2 = (Sum2 - ((i*d%maxn2+d)/maxn2)%mod*(d2%mod)%mod*(t%maxn2)%mod + mod) % mod
		}

		t4 = t4 * maxn2 % mod

		p := make([]int, 2000200)
		for i := 0; i < maxn2; i++ {
			p[i] = ((t/maxn2)%mod*Sum + t4) % mod
			p[i] = (p[i] + h[i] + (t/maxn2)%mod*d%mod*d2%mod*(t%maxn2)) % mod
			Sum = (Sum + d2) % mod
		}

		Ans := 0
		for i := 1; i <= maxn1; i++ {
			t2 := Find(i, X)
			if i <= m%maxn1 {
				t3 := p[t2%maxn2]
				t3 = (t3 + Find(t2%maxn2+t*d, Y)) % mod
				t3 = (t3 + ((t+1)%mod)*((t2/maxn2)%mod*d2%mod)) % mod
				Ans = (Ans + t3) % mod
			} else {
				t3 := p[t2%maxn2]
				t3 = (t3 + (t%mod)*((t2/maxn2)%mod*d2%mod)) % mod
				Ans = (Ans + t3) % mod
			}
		}
		fmt.Println(Ans)
	} else {
		Ans := 0
		for i := 1; i <= m; i++ {
			Ans = (Ans + Find(i, K)) % mod
		}
		fmt.Println(Ans)
	}
}
