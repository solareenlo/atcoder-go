package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var w, h, A, R int
	fmt.Scan(&w, &h, &R, &A)
	R--

	xs := make([]int, A)
	rs := make([]int, A)
	for i := 0; i < A; i++ {
		var r, B int
		fmt.Scan(&r, &B)
		r--
		x := 0
		rs[i] = r
		for j := 0; j < B; j++ {
			var k int
			fmt.Scan(&k)
			x |= 1 << (k - 1)
		}
		xs[i] = x
	}

	if h == 1 {
		{
			y := 0
			if A > 0 {
				y = xs[0]
			}
			push := 0
			for i := 0; i < w; i++ {
				if (y<<((^i)&0x3f)) < 0 && i+1 < w {
					push |= 1 << (i + 1)
					y ^= 7 << i
					y &= (1 << w) - 1
				}
			}
			if y == 0 {
				PRINT(push)
				return
			}
		}
		{
			y := 0
			if A > 0 {
				y = xs[0]
			}
			push := 0
			push |= 1 << 0
			y ^= 3
			y &= (1 << w) - 1
			for i := 0; i < w; i++ {
				if (y<<((^i)&0x3f)) < 0 && i+1 < w {
					push |= 1 << (i + 1)
					y ^= 7 << i
					y &= (1 << w) - 1
				}
			}
			if y == 0 {
				PRINT(push)
				return
			}
		}
		return
	}

	M := make([][]bool, w*2)
	for i := range M {
		M[i] = make([]bool, w*2)
	}
	for i := 0; i < w; i++ {
		M[i][i+w] = true
		M[i+w][i] = true
		for j := -1; j <= 1; j++ {
			if i+j >= 0 && i+j < w {
				M[i+j][i] = true
			}
		}
	}
	ps := Ps(M, 61)

	E := make([][]bool, w)
	for i := range E {
		E[i] = make([]bool, w)
	}

	for i := 0; i < w; i++ {
		v := make([]bool, w*2)
		if i-1 >= 0 {
			v[i-1] = true
		}
		v[i] = true
		if i+1 < w {
			v[i+1] = true
		}
		v[i+w] = true
		res := POW(ps, v, h-1)
		for j := 0; j < w; j++ {
			E[j][i] = res[j]
		}
	}

	to := make([]bool, w)
	for i := 0; i < A; i++ {
		v := make([]bool, w*2)
		for j := 0; j < w; j++ {
			if (xs[i] << ((^j) & 0x3f)) < 0 {
				v[j] = true
			}
		}
		res := POW(ps, v, h-1-rs[i])
		for j := 0; j < w; j++ {
			to[j] = (to[j] && !res[j]) || (!to[j] && res[j])
		}
	}

	result := gaussElimination(E, to)
	head := result.Sol

	ret := make([]bool, w)
	for i := 0; i < A; i++ {
		if rs[i] < R {
			v := make([]bool, w*2)
			for j := 0; j < w; j++ {
				if (xs[i] << ((^j) & 0x3f)) < 0 {
					v[j] = true
				}
			}
			res := POW(ps, v, R-1-rs[i])
			for j := 0; j < w; j++ {
				ret[j] = (ret[j] && !res[j]) || (!ret[j] && res[j])
			}
		}
	}
	for i := 0; i < w; i++ {
		if head[i] && R-1 >= 0 {
			v := make([]bool, w*2)
			if i-1 >= 0 {
				v[i-1] = true
			}
			v[i] = true
			if i+1 < w {
				v[i+1] = true
			}
			v[i+w] = true
			res := POW(ps, v, R-1)
			for j := 0; j < w; j++ {
				ret[j] = (ret[j] && !res[j]) || (!ret[j] && res[j])
			}
		}
	}
	if R == 0 {
		for j := 0; j < w; j++ {
			ret[j] = (ret[j] && !head[j]) || (!ret[j] && head[j])
		}
	}

	lret := 0
	for j := 0; j < w; j++ {
		if ret[j] {
			lret |= 1 << j
		}
	}
	PRINT(lret)
}

func PRINT(push int) {
	fmt.Printf("%v", bits.OnesCount64(uint64(push)))
	for i := 0; i < 61; i++ {
		if (push << uint((^i)&0x3f)) < 0 {
			fmt.Printf(" %v", i+1)
		}
	}
	fmt.Printf("\n")
}

func Mul(A [][]bool, v []bool) []bool {
	m := len(A)
	n := len(v)
	w := make([]bool, m)
	for i := 0; i < m; i++ {
		var sum bool
		for k := 0; k < n; k++ {
			tmp := A[i][k] && v[k]
			if (sum && !tmp) || (!sum && tmp) {
				sum = true
			} else {
				sum = false
			}
		}
		w[i] = sum
	}
	return w
}

func p2(A [][]bool) [][]bool {
	n := len(A)
	C := make([][]bool, n)
	for i := 0; i < n; i++ {
		C[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			var sum bool
			for k := 0; k < n; k++ {
				tmp := A[i][k] && A[k][j]
				if (sum && !tmp) || (!sum && tmp) {
					sum = true
				} else {
					sum = false
				}
			}
			C[i][j] = sum
		}
	}
	return C
}

func Ps(A [][]bool, m int) [][][]bool {
	ret := make([][][]bool, m+1)
	ret[0] = A
	for i := 1; i <= m; i++ {
		ret[i] = p2(ret[i-1])
	}
	return ret
}

func POW(ps [][][]bool, v []bool, e int) []bool {
	d := 0
	for ; e > 0; e >>= 1 {
		if e&1 == 1 {
			v = Mul(ps[d], v)
		}
		d++
	}
	return v
}

type Result struct {
	Mat    [][]bool
	Sol    []bool
	Rank   int
	Exists bool
}

func gaussElimination(M [][]bool, v []bool) Result {
	n := len(M)
	m := len(M[0])
	head := make([]int, n)
	row := 0
	for col := 0; col < m; col++ {
		pivotFound := false
		for prow := row; prow < n; prow++ {
			if M[prow][col] {
				if prow != row {
					for k := 0; k < m; k++ {
						u := M[prow][k]
						M[prow][k] = M[row][k]
						M[row][k] = u
					}
					dum := v[prow]
					v[prow] = v[row]
					v[row] = dum
				}
				pivotFound = true
				break
			}
		}
		if !pivotFound {
			continue
		}
		head[row] = col
		for j := row + 1; j < n; j++ {
			if M[j][col] {
				for k := col; k < m; k++ {
					M[j][k] = (M[j][k] && !M[row][k]) || (!M[j][k] && M[row][k])
				}
				v[j] = (v[j] && !v[row]) || (!v[j] && v[row])
			}
		}
		row++
	}
	ret := Result{}
	ret.Mat = M
	for i := row; i < n; i++ {
		if v[i] {
			ret.Rank = row
			ret.Exists = false
			return ret
		}
	}
	for i := row - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if M[j][head[i]] {
				for k := head[i]; k < m; k++ {
					M[j][k] = (M[j][k] && !M[i][k]) || (!M[j][k] && M[i][k])
				}
				v[j] = (v[j] && !v[i]) || (!v[j] && v[i])
			}
		}
	}
	retv := make([]bool, m)
	for i := 0; i < row; i++ {
		retv[head[i]] = v[i]
	}
	ret.Sol = retv
	ret.Rank = row
	ret.Exists = true
	return ret
}
