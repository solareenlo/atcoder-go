package main

import (
	"bufio"
	"fmt"
	"os"
)

const AS = 200010

var dat = [2 * AS]int{}

func modify(p, value int) {
	p += AS
	for dat[p] = value; p > 1; p >>= 1 {
		dat[p>>1] = max(dat[p], dat[p^1])
	}
}

func query(l, r int) int {
	ret := 0
	l += AS
	r += AS
	for ; l < r; l, r = l>>1, r>>1 {
		if l&1 != 0 {
			ret = max(ret, dat[l])
			l++
		}
		if r&1 != 0 {
			r--
			ret = max(ret, dat[r])
		}
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	A := [200010]int{}
	for i := 1; i < N+1; i++ {
		A[i] = i
	}
	A[0] = 0
	A[N+1] = 1e9

	for i := 0; i < Q; i++ {
		var t, x, y int
		fmt.Fscan(in, &t, &x, &y)
		if t == 1 {
			A[x], A[x+1] = A[x+1], A[x]
			if A[x-1] > A[x] && A[x-1] < A[x+1] {
				modify(x-1, x-1)
			} else if A[x-1] < A[x] && A[x-1] > A[x+1] {
				modify(x-1, 0)
			}
			if A[x] > A[x+1] {
				modify(x, x)
			} else {
				modify(x, 0)
			}
			if A[x+1] > A[x+2] && A[x] < A[x+2] {
				modify(x+1, x+1)
			} else if A[x+1] < A[x+2] && A[x] > A[x+2] {
				modify(x+1, 0)
			}
		} else {
			k := query(x, y)
			for k > 0 {
				A[k], A[k+1] = A[k+1], A[k]
				if A[k-1] > A[k] && A[k-1] < A[k+1] {
					modify(k-1, k-1)
				}
				modify(k, 0)
				if A[k+1] > A[k+2] && A[k] < A[k+2] {
					modify(k+1, k+1)
				}
				k = query(x, y)
			}
		}
	}

	for i := 1; i < N; i++ {
		fmt.Fprint(out, A[i], " ")
	}
	fmt.Fprintln(out, A[N])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
