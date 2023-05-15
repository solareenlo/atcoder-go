package main

import "fmt"

func main() {
	var L, R, M int
	fmt.Scan(&L, &R, &M)
	fmt.Println(f(R, M) - f(L-1, M))
}

func f(N, M int) int {
	ret := N * 5
	p := 1
	for k := 1; k < M; k++ {
		p *= 10
		ord := 1 << (k - 1)
		n := N
		t := 1
		for t*5 < p {
			t *= 5
			n--
		}
		if n < 0 {
			continue
		}
		for n%ord != 0 {
			n--
			t = t * 5 % (p * 10)
			ret += t / p
		}
		x := 0
		for i := 0; i < ord; i++ {
			t = t * 5 % (p * 10)
			x += t / p
		}
		ret += n / ord * x
	}
	return ret
}
