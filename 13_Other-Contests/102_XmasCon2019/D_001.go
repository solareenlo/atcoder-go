package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	v := int(math.Sqrt(float64(N) + 0.5))
	n_6 := int(math.Cbrt(float64(v) + 0.5))
	K := int(math.Pow(float64(N)+0.5, 2.0/3.0))
	B := N / K

	primes := make([]int, 0)
	is_prime := make([]bool, v+1)
	primes = append(primes, 2)
	for i := 3; i*i <= v; i += 2 {
		for j := i * i; !is_prime[i] && j <= v; j += (i << 1) {
			is_prime[j] = true
		}
	}
	for i := 3; i <= v; i += 2 {
		if !is_prime[i] {
			primes = append(primes, i)
		}
	}
	primes = append(primes, v+1)

	pi_6 := 0
	for primes[pi_6] <= n_6 {
		pi_6++
	}

	s0 := make([]int, v+1)
	l0 := make([]int, v+2)
	var divide func(int, int) int
	divide = func(n, d int) int { return int(float64(n) / float64(d)) }
	for i := 1; i <= v; i++ {
		s0[i] = i
		l0[i] = divide(N, i)
	}
	for id := 0; id < pi_6; id++ {
		p := primes[id]
		t := divide(v, p)
		M := divide(N, p)
		for i := 1; i <= t; i++ {
			l0[i] -= l0[i*p]
		}
		for i := t + 1; i <= v; i++ {
			l0[i] -= s0[divide(M, i)]
		}
		for i, j := v, t; j > 0; j-- {
			for e := j * p; i >= e; i-- {
				s0[i] -= s0[j]
			}
		}
	}

	sf := make([]int, v+1)
	lf := make([]int, v+2)

	_v := int(math.Sqrt(float64(K) + 0.5))
	sf[1] = 2
	for i := pi_6; primes[i] <= _v; i++ {
		t := divide(v, primes[i])
		m := divide(K, primes[i])
		M := divide(N, primes[i])
		j := i + 1
		for ; primes[j] <= t; j++ {
			sf[primes[i]*primes[j]] += 2
		}
		for ; primes[j] <= m; j++ {
			lf[divide(M, primes[j])] += 2
		}
		if primes[i]*primes[i] <= v {
			sf[primes[i]*primes[i]] += 1
		} else {
			lf[divide(M, primes[i])] += 1
		}
	}

	_v = int(math.Sqrt(float64(v) + 0.5))
	for i := pi_6; primes[i] <= _v; i++ {
		q := primes[i] * primes[i]
		t := divide(v, q)
		m := divide(K, q)
		M := divide(N, q)
		j := pi_6
		for ; primes[j] <= t; j++ {
			sf[q*primes[j]] += 1
		}
		for ; primes[j] <= m; j++ {
			lf[divide(M, primes[j])] += 1
		}
	}

	for i := 1; i <= v; i++ {
		sf[i] += sf[i-1]
	}
	lf[v] += sf[v]
	for i := v - 1; i > B; i-- {
		lf[i] += lf[i+1]
	}
	for i := 1; i <= v; i++ {
		sf[i] -= s0[i]
		lf[i] -= l0[i]
	}

	roughs := make([]int, 0)
	for i := n_6 + 1; i <= v; i++ {
		if s0[i] != s0[i-1] {
			roughs = append(roughs, i)
		}
	}
	roughs = append(roughs, v+1)
	for i := B; i >= 1; i-- {
		lf[i] = 1 - l0[i]
		m := divide(N, i)
		u := int(math.Sqrt(float64(m) + 0.5))
		t := divide(v, i)
		k := 0
		for ; roughs[k] <= t; k++ {
			lf[i] -= lf[i*roughs[k]] + (sf[roughs[k]]-sf[roughs[k]-1])*l0[i*roughs[k]]
		}
		for now := divide(m, roughs[k]); roughs[k] <= u; {
			lf[i] -= sf[now] + (sf[roughs[k]]-sf[roughs[k]-1])*s0[now]
			k++
			now = divide(m, roughs[k])
		}
		lf[i] += s0[u] * sf[u]
	}

	for id := pi_6 - 1; id >= 0; id-- {
		p := primes[id]
		t := divide(v, p)
		M := divide(N, p)
		for i := 1; i <= t; i++ {
			lf[i] -= lf[i*p]
		}
		for i := t + 1; i <= v; i++ {
			lf[i] -= sf[divide(M, i)]
		}
		for i, j := v, t; j > 0; j-- {
			for e := j * p; i >= e; i-- {
				sf[i] -= sf[j]
			}
		}
	}

	ans := 0
	for i := 1; i*i <= v; i++ {
		ans += lf[i*i]
	}
	for i := int(math.Sqrt(float64(v)+0.5)) + 1; i <= v; i++ {
		ans += sf[N/(i*i)]
	}
	fmt.Println(ans)
}
