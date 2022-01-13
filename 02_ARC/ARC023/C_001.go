package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	cnt := 1
	c := 0
	pre := a[0]
	for i := 1; i < n; i++ {
		if a[i] == -1 {
			c++
		} else {
			dif := a[i] - pre
			for k := 0; k < c; k++ {
				cnt *= dif + c - k
				cnt %= mod
				cnt = divMod(cnt, k+1)
			}
			pre = a[i]
			c = 0
		}
	}

	fmt.Println(cnt)
}

const mod = 1000000007

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
