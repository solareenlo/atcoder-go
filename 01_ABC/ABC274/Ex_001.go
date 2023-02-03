package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 500500
	const Base = 124136381

	var n, q int
	fmt.Fscan(in, &n, &q)

	var pw [N]int
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = mul(pw[i-1], Base)
	}

	var A, hs [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i])
		hs[i] = mul(hs[i-1], Base) ^ A[i]
	}
	var getHs func(int, int) int
	getHs = func(l, r int) int {
		return hs[r] ^ mul(hs[l-1], pw[r-l+1])
	}
	for q > 0 {
		q--
		var a, b, c, d, e, f int
		fmt.Fscan(in, &a, &b, &c, &d, &e, &f)
		l, r := 0, min(b-a+1, f-e+1)
		var chk func(int) bool
		chk = func(len int) bool {
			return (getHs(a, a+len-1) ^ getHs(c, c+len-1)) == getHs(e, e+len-1)
		}
		for l < r {
			mid := (l + r + 1) / 2
			if chk(mid) {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if l == f-e+1 {
			fmt.Fprintln(out, "No")
			continue
		}
		if l == b-a+1 {
			fmt.Fprintln(out, "Yes")
			continue
		}
		if (A[a+l] ^ A[c+l]) < A[e+l] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func mul(a, b int) int {
	ans := 0
	for i := 0; i < 61; i++ {
		if (b>>i)&1 != 0 {
			ans ^= shift(a, i)
		}
	}
	return ans
}

func shift(a, k int) int {
	return ((a & ((1 << (61 - k)) - 1)) << k) | (a >> (61 - k))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
