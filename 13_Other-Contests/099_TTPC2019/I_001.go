package main

import "fmt"

func main() {
	var P, Q, L, R int
	fmt.Scan(&P, &Q, &L, &R)
	mod = Q

	if Q%P != 0 {
		ans := 1
		for i := L; i <= R && ans != 0; i++ {
			t := i
			for t%P == 0 {
				t /= P
			}
			ans = t % Q * ans % Q
		}
		fmt.Println(ans)
		return
	}

	var FL [10000000]int
	FL[0] = 1
	for i := 1; i < Q; i++ {
		if i%P != 0 {
			FL[i] = FL[i-1] * i % Q
		} else {
			FL[i] = FL[i-1] * 1 % Q
		}
	}
	var FR [10000000]int
	FR[Q-1] = Q - 1
	for i := Q - 1; i > 0; {
		i--
		if i%P != 0 {
			FR[i] = FR[i+1] * i % Q
		} else {
			FR[i] = FR[i+1] * 1 % Q
		}
	}
	FR[0] = 1
	ans := 1
	for R-L+1 >= Q {
		l := (L + Q - 1) / Q * Q
		r := R / Q * Q
		ans = ans * FR[L%Q] % Q * FL[R%Q] % Q * powMod(FL[Q-1], (r-l)/Q) % Q
		L = (L + P - 1) / P
		R = R / P
	}
	for i := L; i <= R; i++ {
		t := i
		for t%P == 0 {
			t /= P
		}
		ans = t % Q * ans % Q
	}
	fmt.Println(ans)
}

var mod int

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
