package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{ x, y int }

var (
	L, i, s, R, j, t int
)

func lift(rle []pair, X, p int) []pair {
	L, i, s, R, j, t = 0, 0, 0, 0, 0, 0
	asc_l := func(x int) {
		L += x
		for x != 0 {
			rest := rle[i].y - s
			if x < rest {
				break
			}
			x -= rest
			i++
			s = 0
		}
		s += x
	}
	asc_r := func(x int) {
		R += x
		for x != 0 {
			rest := rle[j].y - t
			if x < rest {
				break
			}
			x -= rest
			j++
			t = 0
		}
		t += x
	}

	var res []pair
	asc_r(X)
	for j != len(rle) {
		val := rle[i].x + rle[j].x
		if val >= p {
			val -= p
		}
		len := min(rle[i].y-s, rle[j].y-t)
		res = append(res, pair{val, len})
		asc_l(len)
		asc_r(len)
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)

	rle := make([]pair, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &rle[i].x, &rle[i].y)
	}

	for n, x := N, 1; n > K; n -= x {
		x = 1
		for x*7 <= n-K {
			x *= 7
		}
		rle = lift(rle, x, 7)
	}

	decode := func(r []pair) []int {
		d := make([]int, 0)
		for i := range r {
			k := r[i].x
			v := r[i].y
			for i := 0; i < v; i++ {
				d = append(d, k)
			}
		}
		return d
	}

	res := decode(rle)
	for i := 0; i < K; i++ {
		fmt.Fprint(out, res[i], " ")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
