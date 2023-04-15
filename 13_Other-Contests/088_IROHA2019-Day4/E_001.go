package main

import (
	"fmt"
	"math/bits"
)

const cys = 1000000007

var fac, inv [2005]int

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fac[0] = 1
	inv[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % cys
	}
	inv[n] = qpow(fac[n], cys-2)
	for i := n - 1; i >= 1; i-- {
		inv[i] = inv[i+1] * (i + 1) % cys
	}
	fmt.Println(mod(qpow((m+1)%cys, n) + cys - getans(n, m)))
}

func getans(n, m int) int {
	if m == 0 {
		return 1
	}
	t := 1 << (63 - countLeadingZeros(uint64(m)))
	ret := 0
	for i := 0; i < n; i += 2 {
		ret = mod(ret + qpow((m-t+1)%cys, i)*qpow(t%cys, n-i-1)%cys*fac[n]%cys*inv[i]%cys*inv[n-i]%cys)
	}
	if (n & 1) == 0 {
		ret = mod(ret + getans(n, m-t))
	}
	return ret
}

func countLeadingZeros(x uint64) int {
	return bits.LeadingZeros64(x)
}

func mod(x int) int {
	if x >= cys {
		return x - cys
	}
	return x
}

func qpow(x, p int) int {
	ret := 1
	for ; p > 0; p >>= 1 {
		if (p & 1) != 0 {
			ret = ret * x % cys
		}
		x = x * x % cys
	}
	return ret
}
