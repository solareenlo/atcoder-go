package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

type pair struct {
	x, y int
}

var seg, tar [23][524293]int
var pw [524293]int

func change(a, b, l, r, k, x, px, y, py int) {
	if r <= a || b <= l {
		return
	}
	if a <= l && r <= b {
		tar[px][k], tar[py][k] = tar[py][k], tar[px][k]
		return
	}
	change(a, b, l, (l+r)>>1, k*2, tar[x][k*2], x, tar[y][k*2], y)
	change(a, b, (l+r)>>1, r, k*2+1, tar[x][k*2+1], x, tar[y][k*2+1], y)
	seg[x][k] = (seg[tar[x][k*2]][k*2] + seg[tar[x][k*2+1]][k*2+1]*pw[(r-l)>>1]) % MOD
	seg[y][k] = (seg[tar[y][k*2]][k*2] + seg[tar[y][k*2+1]][k*2+1]*pw[(r-l)>>1]) % MOD
}

func calc(a, b, l, r, k, x int) pair {
	if r <= a || b <= l {
		return pair{0, 0}
	}
	if a <= l && r <= b {
		return pair{seg[x][k], r - l}
	}
	lc := calc(a, b, l, (l+r)>>1, k*2, tar[x][k*2])
	rc := calc(a, b, (l+r)>>1, r, k*2+1, tar[x][k*2+1])
	return pair{(lc.x + rc.x*pw[lc.y]) % MOD, lc.y + rc.y}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)
	max_size := 2
	for max_size < N {
		max_size *= 2
	}
	pw[0] = 1
	for i := 1; i <= max_size; i++ {
		pw[i] = 1000000 * pw[i-1] % MOD
	}
	for i := 0; i < M; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < N; j++ {
			seg[i][max_size+j] = int(s[j] - 96)
		}
		for j := max_size / 2; j >= 1; j >>= 1 {
			mul := pw[max_size/(j*2)]
			for k := j; k < j*2; k++ {
				seg[i][k] = (seg[i][k*2] + seg[i][k*2+1]*mul) % MOD
			}
		}
		for j := 0; j < max_size*2; j++ {
			tar[i][j] = i
		}
	}
	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var tp, x, y, l, r int
		fmt.Fscan(in, &tp, &x, &y, &l, &r)
		x--
		y--
		l--
		if tp == 1 {
			change(l, r, 0, max_size, 1, x, -1, y, -1)
		} else {
			res := calc(l, r, 0, max_size, 1, x)
			fmt.Fprintln(out, res.x)
		}
	}
}
