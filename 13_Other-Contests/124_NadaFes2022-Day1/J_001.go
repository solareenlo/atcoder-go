package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var insu [1000001]int
	S := 1000000
	for i := 2; i <= S; i++ {
		if insu[i] != 0 {
			continue
		}
		for j := i; j <= S; j += i {
			insu[j] = i
		}
	}
	var aru [1000001]int
	all_ans := 1
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
		x := A[i]
		for x > 1 {
			y := insu[x]
			z := 1
			for y == insu[x] {
				z *= y
				aru[z]++
				x /= y
			}
		}
	}
	for i := 2; i <= S; i++ {
		if aru[i] > 0 {
			all_ans = all_ans * insu[i] % MOD
		}
	}
	var Q int
	fmt.Fscan(in, &Q)
	var kesu [1000001]int
	var kepa [1000001]int
	for kai := 1; kai <= Q; kai++ {
		waru := 1
		var K int
		fmt.Fscan(in, &K)
		for K > 0 {
			K--
			var xx int
			fmt.Fscan(in, &xx)
			x := A[xx-1]
			for x > 1 {
				y := insu[x]
				z := 1
				for y == insu[x] {
					z *= y
					if kepa[z] != kai {
						kepa[z] = kai
						kesu[z] = 1
					} else {
						kesu[z]++
					}
					if kesu[z] == aru[z] {
						waru = waru * insu[z] % MOD
					}
					x /= y
				}
			}
		}
		fmt.Println(divMod(all_ans, waru))
	}
}

const MOD = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a int) int {
	b, u, v := MOD, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
