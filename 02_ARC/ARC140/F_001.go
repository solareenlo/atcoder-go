package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 250000
const L = 19 /* L = ceil(log2(N * 2)) */
const N_ = 525000
const MD = 998244353

var vv, ff, gg, pp2 []int
var vv_ []int
var wwu, wwv [][]int

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Init()

	var n, m int
	fmt.Scan(&n, &m)
	r := n % m
	l_ := 0
	for 1<<l_ < n {
		l_++
	}

	var bb []int
	bb = make([]int, N_)
	n_ := 1 << l_
	for i := 0; i < n_; i++ {
		bb[i] = 1
	}
	var aa []int
	aa = make([]int, N_)
	solve(aa, n/m)
	ntt(aa, l_, 0)
	for i := 0; i < n_; i++ {
		bb[i] = bb[i] * power(aa[i], m-r) % MD
	}
	if r != 0 {
		for i := range aa {
			aa[i] = 0
		}
		solve(aa, n/m+1)
		ntt(aa, l_, 0)
		for i := 0; i < n_; i++ {
			bb[i] = bb[i] * power(aa[i], r) % MD
		}
	}
	ntt(bb, l_, 1)
	for i := 0; i < n; i++ {
		bb[i] = bb[i] * vv_[l_] % MD
	}
	for i := range aa {
		aa[i] = 0
	}
	for i := 0; i <= n-m; i++ {
		aa[i] = bb[n-m-i] * ff[n-i] % MD * ff[i] % MD
	}
	for i := range bb {
		bb[i] = 0
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			bb[n-i] = gg[i]
		} else {
			bb[n-i] = MD - gg[i]
		}
	}
	ntt(aa, l_+1, 0)
	ntt(bb, l_+1, 0)
	for i := 0; i < n_*2; i++ {
		aa[i] = aa[i] * bb[i] % MD
	}
	ntt(aa, l_+1, 1)
	for i := 0; i < n_*2; i++ {
		aa[i] = aa[i] * vv_[l_+1] % MD
	}
	for i := 0; i < n; i++ {
		aa[i] = aa[n+i] * gg[i] % MD
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", aa[i])
	}
	fmt.Fprintln(out)
}

func ntt_(aa []int, l, inverse int) {
	if l != 0 {
		n := 1 << l
		m := n >> 1
		var ww []int
		if inverse != 0 {
			ww = make([]int, len(wwv[l]))
			copy(ww, wwv[l])
		} else {
			ww = make([]int, len(wwu[l]))
			copy(ww, wwu[l])
		}
		ntt_(aa, l-1, inverse)
		tmp := aa[m:]
		ntt_(tmp, l-1, inverse)
		for i := 0; i+m < n; i++ {
			j := i + m
			a := aa[i]
			b := aa[j] * ww[i] % MD
			aa[i] = a + b
			if aa[i] >= MD {
				aa[i] -= MD
			}
			aa[j] = a - b
			if aa[j] < 0 {
				aa[j] += MD
			}
			j = i + m
		}
	}
}

func ntt(aa []int, l, inverse int) {
	n := 1 << l
	for i, j := 0, 1; j < n; j++ {
		m := n >> 1
		i ^= m
		for i < m {
			m >>= 1
			i ^= m
		}
		if i < j {
			aa[i], aa[j] = aa[j], aa[i]
		}
	}
	ntt_(aa, l, inverse)
}

func solve_(aa [][][]int, n int) {
	aa_ := make([][][]int, 2)
	bb_ := make([][][]int, 2)
	for i := range aa_ {
		aa_[i] = make([][]int, 2)
		bb_[i] = make([][]int, 2)
		for j := range aa_[i] {
			aa_[i][j] = make([]int, N_)
			bb_[i][j] = make([]int, N_)
		}
	}
	if n == 1 {
		for w := 0; w < 2; w++ {
			for x := 0; x < 2; x++ {
				if w != 0 && x != 0 {
					aa[w][x][0] = vv[2]
				} else {
					aa[w][x][0] = 1
				}
			}
		}
		return
	}
	m := n / 2
	solve_(aa, m)
	l_ := 0
	for 1<<l_ < n {
		l_++
	}
	n_ := 1 << l_
	for w := 0; w < 2; w++ {
		for x := 0; x < 2; x++ {
			for i := 0; i < n_; i++ {
				if i < m {
					aa_[w][x][i] = aa[w][x][i]
				} else {
					aa_[w][x][i] = 0
				}
			}
			ntt(aa_[w][x], l_, 0)
			for i := 0; i < n; i++ {
				bb_[w][x][i] = 0
			}
		}
	}
	for w := 0; w < 2; w++ {
		for x := 0; x < 2; x++ {
			for y := 0; y < 2; y++ {
				for i := 0; i < n_; i++ {
					var tmp int
					if x == 0 {
						tmp = 1
					} else {
						tmp = wwu[l_][i]
					}
					bb_[w][y][i] = (bb_[w][y][i] + aa_[w][x][i]*aa_[x][y][i]%MD*tmp) % MD
				}
			}
		}
	}
	for w := 0; w < 2; w++ {
		for x := 0; x < 2; x++ {
			ntt(bb_[w][x], l_, 1)
			for i := 0; i < n_; i++ {
				bb_[w][x][i] = bb_[w][x][i] * vv_[l_] % MD
			}
			for i := 0; i < n; i++ {
				aa[w][x][i] = 0
			}
		}
	}
	for w := 0; w < 2; w++ {
		for x := 0; x < 2; x++ {
			if n%2 == 0 {
				for i := 0; i < n; i++ {
					aa[w][x][i] = bb_[w][x][i]
				}
			} else {
				for y := 0; y < 2; y++ {
					for i := 0; i < n; i++ {
						var tmp int
						if (x & y) != 0 {
							tmp = vv[2]
						} else {
							tmp = 1
						}
						aa[w][y][i+x] = (aa[w][y][i+x] + bb_[w][x][i]*tmp) % MD
					}
				}
			}
		}
	}
}

func solve(aa []int, n int) {
	aa_ := make([][][]int, 2)
	for i := range aa_ {
		aa_[i] = make([][]int, 2)
		for j := range aa_[i] {
			aa_[i][j] = make([]int, N_)
		}
	}
	solve_(aa_, n)
	for i := 0; i < n; i++ {
		aa[i] = aa_[1][1][i] * pp2[i+1] % MD
	}
}

func Init() {
	vv = make([]int, N+1)
	ff = make([]int, N+1)
	gg = make([]int, N+1)
	pp2 = make([]int, N+1)
	vv_ = make([]int, L+1)
	wwu = make([][]int, L+1)
	wwv = make([][]int, L+1)
	ff[0] = 1
	gg[0] = 1
	pp2[0] = 1
	for i := 1; i <= N; i++ {
		if i == 1 {
			vv[i] = 1
		} else {
			vv[i] = vv[i-MD%i] * (MD/i + 1) % MD
		}
		ff[i] = ff[i-1] * i % MD
		gg[i] = gg[i-1] * vv[i] % MD
		pp2[i] = pp2[i-1] * 2 % MD
	}
	u := power(3, (MD-1)>>L)
	v := power(u, MD-2)
	for l := L; l > 0; l-- {
		n := 1 << l
		vv_[l] = power(1<<l, MD-2)
		wwu[l] = make([]int, n)
		wwv[l] = make([]int, n)
		wwu[l][0] = 1
		wwv[l][0] = 1
		for i := 1; i < n; i++ {
			wwu[l][i] = wwu[l][i-1] * u % MD
			wwv[l][i] = wwv[l][i-1] * v % MD
		}
		u = u * u % MD
		v = v * v % MD
	}
	vv_[0] = 1
}

func power(a, k int) int {
	p := 1
	for k != 0 {
		if (k & 1) > 0 {
			p = p * a % MD
		}
		a = a * a % MD
		k >>= 1
	}
	return p
}
