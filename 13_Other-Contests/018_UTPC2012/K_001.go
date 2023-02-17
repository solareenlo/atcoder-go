package main

import (
	"fmt"
	"math"
)

var a, b int

func main() {
	x := []int{5, 3, 2, 0, 4, 1}
	y := []int{4, 2, 0, 3, 1, 5}
	fmt.Scan(&a, &b)
	if a == 0 || b == 0 {
		fmt.Println(4.0)
		return
	}
	g := gcd(a, b)
	a /= g
	b /= g
	z := sol(a, b, x, y)
	dl := zyb(z)
	da := float64(a)
	db := float64(b)
	ans := float64(dl) * math.Sqrt(da*da+db*db)
	fmt.Println(ans)
}

func sol(a, b int, x, y []int) []int {
	if b == 0 {
		return x
	}
	if a == 0 {
		return y
	}
	if a >= b {
		t1 := qp(x, a/b)
		t2 := mul(t1, y)
		return sol(a%b, b, x, t2)
	} else {
		t1 := qp(y, b/a)
		t2 := mul(t1, x)
		return sol(a, b%a, t2, y)
	}
}

func qp(a []int, b int) []int {
	ans := make([]int, 0)
	for i := 0; i < 6; i++ {
		ans = append(ans, i)
	}
	for b != 0 {
		if (b & 1) != 0 {
			ans = mul(ans, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return ans
}

func zyb(x []int) int {
	y := make([]int, len(x))
	for i := range x {
		y[i] = x[i]
	}
	s := 1
	for {
		fl := 1
		for i := 0; i < 6; i++ {
			if y[i] != i {
				fl = 0
			}
		}
		if fl != 0 {
			return s
		}
		s++
		y = mul(y, x)
	}
}

func mul(l, r []int) []int {
	s := make([]int, 0)
	for i := 0; i < 6; i++ {
		s = append(s, r[l[i]])
	}
	return s
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
