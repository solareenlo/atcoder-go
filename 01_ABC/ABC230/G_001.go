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

	p := make([]int, n+1)
	mu := make([]int, n+1)
	prime := make([]bool, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &p[i])
		mu[i] = -1
		prime[i] = true
	}

	v := make([]int, 0)
	fac := make([][]int, 200002)
	for i := 2; i < n+1; i++ {
		if prime[i] {
			mu[i] = -mu[i]
			for j := 2; i*j < n+1; j++ {
				prime[i*j] = false
				if j%i != 0 {
					mu[i*j] *= -1
				} else {
					mu[i*j] = 0
				}
			}
		}
		if mu[i] != 0 {
			v = append(v, i)
			for j := i; j < n+1; j += i {
				fac[j] = append(fac[j], i)
			}
		}
	}

	res := 0
	cnt := make([]int, 200002)
	use := make([]int, 200002)
	for _, a := range v {
		can := make([]int, 0)
		for i := a; i < n+1; i += a {
			for _, b := range fac[p[i]] {
				cnt[b]++
				if use[b] == 0 {
					use[b] = 1
					can = append(can, b)
				}
			}
		}
		for _, b := range can {
			res += mu[a] * mu[b] * cnt[b] * (cnt[b] + 1) / 2
			cnt[b] = 0
			use[b] = 0
		}
	}
	fmt.Println(res)
}
