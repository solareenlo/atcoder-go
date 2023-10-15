package main

import (
	"fmt"
	"math"
)

func dfsPyth(a, b, c int, f func(int, int, int) bool) {
	if a <= 0 || b <= 0 || c <= 0 {
		return
	}
	if !f(a, b, c) {
		return
	}
	dfsPyth(a-2*b+2*c, 2*a-b+2*c, 2*a-2*b+3*c, f)
	dfsPyth(a+2*b+2*c, 2*a+b+2*c, 2*a+2*b+3*c, f)
	if a != b {
		dfsPyth(-a+2*b+2*c, -2*a+b+2*c, -2*a+2*b+3*c, f)
	}
}

func isPythRoot(a, b, c int) bool {
	aa := a + 2*b - 2*c
	bb := 2*a + b - 2*c
	cc := -2*a - 2*b + 3*c
	return (aa <= 0 && bb <= 0) || aa == 0 || bb == 0 || cc <= 0
}

func genPyth(k int, f func(int, int, int) bool) {
	var checkPythRoot = func(a, b, c int) {
		if a*a+b*b+k == c*c && isPythRoot(a, b, c) && gcd(gcd(a, b), c) == 1 {
			dfsPyth(a, b, c, f)
		}
	}
	if k == 0 {
		checkPythRoot(3, 4, 5)
	} else {
		for d := int(1); d*d <= 4*abs(k); d++ {
			if 4*abs(k)%d == 0 {
				var checkB = func(b int) {
					if (4*k/b+3*b)%4 == 0 {
						a := (4*k/b + 3*b) / 4
						if 1 <= a && a <= b && (2*a+b)%2 == 0 {
							c := (2*a + b) / 2
							checkPythRoot(a, b, c)
						}
					}
				}
				checkB(d)
				if d != 4*abs(k)/d {
					checkB(4 * abs(k) / d)
				}
			}
		}
	}
	for a := int(1); a*a <= 4*k; a++ {
		for b := a; a*(4*b-3*a) <= 4*k; b++ {
			c := int(math.Sqrt(float64(a*a + b*b + k)))
			if 2*a+b != 2*c {
				checkPythRoot(a, b, c)
			}
		}
	}
	for c := int(1); c*c <= 8*(-k); c++ {
		for a := int(1); 2*a*a <= c*c-k; a++ {
			b := int(math.Sqrt(float64(c*c - k - a*a)))
			if 2*a+b != 2*c {
				checkPythRoot(a, b, c)
			}
		}
	}
}

func main() {
	for {
		var L, S int
		_, err := fmt.Scan(&L, &S)
		if err != nil {
			break
		}
		var ans int
		genPyth(S*S, func(a, b, c int) bool {
			if a+b+c > L {
				return false
			}
			if a+b > c {
				ans++
			}
			return true
		})
		fmt.Println(ans)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
