package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 500050

var (
	n int
	f = [N]int{}
)

func add(x int) {
	for ; x <= n; x += x & -x {
		f[x]++
	}
}

func ask(x int) int {
	y := 0
	for ; x > 0; x -= x & -x {
		y += f[x]
	}
	return y
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1_000_000_007

	fmt.Fscan(in, &n)
	J := make([]int, n+1)
	J[0] = 1
	for i := 1; i <= n; i++ {
		J[i] = J[i-1] * i % mod
	}
	a := make([]int, n+1)
	flag := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		flag[a[i]] = true
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if !flag[i] {
			sum += i - 1
			sum %= mod
		}
	}
	s1 := make([]int, n+2)
	for i := 1; i <= n; i++ {
		s1[i] = s1[i-1]
		if !flag[i] {
			s1[i]++
		}
	}
	s2 := make([]int, n+2)
	for i := n; i > 0; i-- {
		s2[i] = s2[i+1]
		if !flag[i] {
			s2[i]++
		}
	}
	m := s1[n]
	dat := 0
	cnt := 0
	ans := 0
	for i := 1; i <= n; i++ {
		anss := 0
		if a[i] != 0 {
			add(a[i])
			anss = (anss + (a[i]-ask(a[i]))*J[m]) % mod
			if m != 0 {
				anss = (anss - s1[a[i]]*cnt%mod*J[m-1]) % mod
			}
			dat += s2[a[i]]
			dat %= mod
		} else {
			anss = (anss + (sum-dat)*J[m-1]) % mod
			if m > 1 {
				anss = (anss - m*(m-1)/2%mod*cnt%mod*J[m-2]) % mod
			}
			cnt++
		}
		ans = (ans + anss*J[n-i]) % mod
	}
	fmt.Println(((ans+J[m])%mod + mod) % mod)
}
