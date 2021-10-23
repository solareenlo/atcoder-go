package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	la, ra := 0, 0
	for l>>la != 0 {
		la++
	}
	la--
	for r>>ra != 0 {
		ra++
	}
	ra--

	mod := int(1e9 + 7)
	res := 0
	for i := la; i <= ra; i++ {
		dp1, dp2, dp3, dp4 := 1, 0, 0, 0
		lc, rc := 0, ^0
		if la == i {
			lc = l
		}
		if ra == i {
			rc = r
		}
		for j := i - 1; j >= 0; j-- {
			lb, rb := lc>>j&1, rc>>j&1
			lb ^= 1
			dp4 = (dp4*3 + rb*dp2 + lb*dp3) % mod
			if rb != 0 {
				dp2 = (dp2*2 + lb*dp1) % mod
			}
			if lb != 0 {
				dp3 = (dp3*2 + rb*dp1) % mod
			}
			dp1 &= lb | rb
		}
		res += (dp1 + dp2 + dp3 + dp4) % mod
		res %= mod
	}

	fmt.Println(res)
}
