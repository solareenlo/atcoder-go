package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

type tuple struct {
	x, y, z int
}

func main() {
	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		solve()
	}
}

func solve() {
	var A, M int
	fmt.Fscan(in, &A, &M)
	if M == 1 {
		fmt.Println(1)
		fmt.Println(0)
		return
	}
	pe := make([]tuple, 0)
	primeFactorization(M, &pe)
	mx := M * phi(M, pe)
	// ar max : 1001
	ar := make([]int, 1)
	mod := 1
	cnt := 1
	flg := true
	for _, tmp := range pe {
		p := tmp.x
		e := tmp.y
		prd := tmp.z
		nw := make([]int, 0)
		t := gcd(mod, p-1)
		if gcd(A, p) != 1 {
			// k \≠ 0 and mod p^e can be anything if it is 0
			for _, v := range ar {
				if len(nw) == 1001 {
					break
				}
				for w := v % t; w < p-1; w += t {
					z, _ := Crt([]int{v, w, 0}, []int{mod, p - 1, prd})
					nw = append(nw, z)
					if len(nw) == 1001 {
						break
					}
				}
			}
		} else {
			flg = false
			for _, v := range ar {
				if len(nw) == 1001 {
					break
				}
				for w := v % t; w < p-1; w += t {
					x := solsub(A, w, p, e)
					z, _ := Crt([]int{v, x}, []int{mod, prd * (p - 1)})
					nw = append(nw, z)
					if len(nw) == 1001 {
						break
					}
				}
			}
		}
		cnt *= (p - 1) / t
		mod *= (p - 1) / t
		mod *= prd
		ar = nw
	}
	cnt *= mx / mod
	if flg == true {
		cnt--
	}
	// Output
	ans := make([]int, 0)
	for i := 0; i < len(ar); i++ {
		for j := ar[i]; j < mx; j += mod {
			if j != 0 {
				ans = append(ans, j)
			}
			if 1000 <= len(ans) {
				fmt.Println(cnt)
				for i := range ans {
					if i == len(ans)-1 {
						fmt.Println(ans[i])
					} else {
						fmt.Printf("%d ", ans[i])
					}
				}
				return
			}
		}
	}
	fmt.Println(cnt)
	for i := range ans {
		if i == len(ans)-1 {
			fmt.Println(ans[i])
		} else {
			fmt.Printf("%d ", ans[i])
		}
	}
}

func primeFactorization(n int, pe *[]tuple) { // Prime factorization
	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		nw := 1
		ex := 0
		for n%i == 0 {
			nw *= i
			ex++
			n /= i
		}
		*pe = append(*pe, tuple{i, ex, nw})
	}
	if n != 1 {
		*pe = append(*pe, tuple{n, 1, n})
	}
}

func phi(M int, pe []tuple) int { // Euler's totient for M
	for _, tmp := range pe {
		p := tmp.x
		M /= p
		M *= (p - 1)
	}
	return M
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func solsub(A, r, p, e int) int {
	res := r
	m := 1
	for i := 1; i <= e; i++ {
		m *= p
		res, _ = Crt([]int{res, powMod2(A, res, m)}, []int{p - 1, m})
	}
	return res % (m * (p - 1))
}

func powMod2(a, n, mod int) int {
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

// Ref: https://github.com/monkukui/ac-library-go
// (rem, mod)
func Crt(r, m []int) (int, int) {
	if len(r) != len(m) {
		panic("")
	}
	n := len(r)
	r0 := int(0)
	m0 := int(1)
	for i := 0; i < n; i++ {
		if !(1 <= m[i]) {
			panic("")
		}
		r1 := SafeMod(r[i], m[i])
		m1 := m[i]
		if m0 < m1 {
			r0, r1 = r1, r0
			m0, m1 = m1, m0
		}
		if m0%m1 == 0 {
			if r0%m1 != r1 {
				return 0, 0
			}
			continue
		}
		g, im := InvGcd(m0, m1)

		u1 := m1 / g
		if (r1-r0)%g != 0 {
			return 0, 0
		}

		x := (r1 - r0) / g % u1 * im % u1

		r0 += x * m0
		m0 *= u1
		if r0 < 0 {
			r0 += m0
		}
	}
	return r0, m0
}

func SafeMod(x, m int) int {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

func InvGcd(a, b int) (int, int) {
	a = SafeMod(a, b)
	if a == 0 {
		return b, 0
	}

	s := b
	t := a
	m0 := int(0)
	m1 := int(1)

	for t > 0 {
		u := s / t
		s -= t * u
		m0 -= m1 * u

		tmp := s
		s = t
		t = tmp
		tmp = m0
		m0 = m1
		m1 = tmp
	}

	if m0 < 0 {
		m0 += b / s
	}
	return s, m0
}
