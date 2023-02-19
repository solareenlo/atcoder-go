package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const mod = 1000000007

var n, ns, sz int
var p, q []int
var v, np, nq, hb []int

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1100100

	var a int
	fmt.Fscan(in, &a)
	n = a
	p = make([]int, N)
	q = make([]int, N)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &p[i], &q[i])
	}
	v = make([]int, N)
	np = make([]int, N)
	nq = make([]int, N)
	for i := 0; i < a; i++ {
		if v[i] != 0 {
			continue
		}
		tmp := 0
		for j := 0; j < a; j++ {
			if p[i] == p[j] {
				tmp += q[j]
				v[j] = 1
			}
		}
		if tmp != 0 {
			np[ns] = p[i]
			nq[ns] = tmp
			ns++
		}
	}
	for i := 0; i < ns; i++ {
		p[i] = np[i]
		q[i] = nq[i]
	}
	a = ns
	n = ns
	deg := 0
	for i := 0; i < a; i++ {
		deg = max(deg, p[i])
	}
	dm := int(10e15)
	for i := 0; i < a; i++ {
		dm = min(dm, p[i])
	}
	if deg <= 1 {
		A, B := 0, 0
		for i := 0; i < a; i++ {
			if p[i] == 1 {
				A += q[i]
			} else {
				B += q[i]
			}
		}
		if A < 0 {
			A = -A
			B = -B
		}
		if A == 0 {
			if B == 0 {
				fmt.Println("Yes 1")
			} else {
				fmt.Println("No")
			}
		} else {
			fmt.Println("No")
		}
		return
	}
	t := 0
	for i := 0; i < a; i++ {
		if p[i] == dm {
			t += q[i]
		}
	}
	t = ABS(t)
	hb = make([]int, N)
	for i := 1; i*i*i <= t; i++ {
		if t%i == 0 {
			if (t/i)%i == 0 {
				hb[sz] = i
				sz++
			}
			s := t / i
			u := int(math.Sqrt(float64(s) + 0.5))
			for j := -100; j <= 100; j++ {
				if u+j > 0 && (u+j)*(u+j) == s {
					hb[sz] = u + j
					sz++
				}
			}
		}
	}
	tmp := hb[:sz]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	for i := 0; i < sz; i++ {
		if calc(hb[i]) == 0 && calc2(hb[i]) == 0 {
			fmt.Printf("Yes %d\n", hb[i])
			return
		}
	}
	fmt.Println("No")
}

func calc(x int) int {
	now := 0
	for i := 0; i < n; i++ {
		now = (now + ppow(x, p[i])*((mod+q[i]%mod)%mod)) % mod
	}
	return now
}

func calc2(x int) int {
	now := 0
	for i := 0; i < n; i++ {
		now = (now + p[i]*ppow(x, p[i]-1)%mod*((mod+q[i]%mod)%mod)) % mod
	}
	return now
}

func ppow(a, b int) int {
	ret := 1
	a %= mod
	for b != 0 {
		if b%2 != 0 {
			ret = ret * a % mod
		}
		b /= 2
		a = a * a % mod
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ABS(a int) int { return max(a, -a) }
