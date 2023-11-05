package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 50
const M = (N * (N - 1) / 2)

var rand_ = rand.New(rand.NewSource(time.Now().UnixNano()))

func cross2(x1, y1, x2, y2 int) int {
	return x1*y2 - x2*y1
}

func cross(x0, y0, x1, y1, x2, y2 int) int {
	return cross2(x1-x0, y1-y0, x2-x0, y2-y0)
}

var xx, yy, zz [N]int

func coplanar(i, j, k, l int) bool {
	x1 := xx[j] - xx[i]
	y1 := yy[j] - yy[i]
	z1 := zz[j] - zz[i]
	x2 := xx[k] - xx[i]
	y2 := yy[k] - yy[i]
	z2 := zz[k] - zz[i]
	x3 := xx[l] - xx[i]
	y3 := yy[l] - yy[i]
	z3 := zz[l] - zz[i]
	return x1*cross2(y2, z2, y3, z3)+y1*cross2(z2, x2, z3, x3)+z1*cross2(x2, y2, x3, y3) == 0
}

func collinear(i, j, k int) bool {
	x1 := xx[i]
	y1 := yy[i]
	z1 := zz[i]
	x2 := xx[j]
	y2 := yy[j]
	z2 := zz[j]
	x3 := xx[k]
	y3 := yy[k]
	z3 := zz[k]
	return cross(x1, y1, x2, y2, x3, y3) == 0 && cross(y1, z1, y2, z2, y3, z3) == 0 && cross(z1, x1, z2, x2, z3, x3) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

var pp, qq, ii, jj [M]int

func compare(h1, h2 int) int {
	if pp[h1] != pp[h2] {
		if pp[h1] < pp[h2] {
			return -1
		}
		return 1
	}
	if qq[h1] != qq[h2] {
		if qq[h1] < qq[h2] {
			return -1
		}
		return 1
	}
	return ii[h1] - ii[h2]
}

func sort(hh []int, l, r int) {
	for l < r {
		i := l
		j := l
		k := r
		h := hh[l+rand_.Int()%(r-l)]
		var tmp int
		for j < k {
			c := compare(hh[j], h)
			if c == 0 {
				j++
			} else if c < 0 {
				tmp = hh[i]
				hh[i] = hh[j]
				hh[j] = tmp
				i++
				j++
			} else {
				k--
				tmp = hh[j]
				hh[j] = hh[k]
				hh[k] = tmp
			}
		}
		sort(hh, l, i)
		l = k
	}
}

func main() {
	hh := make([]int, M)

	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println(1)
		return
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&xx[i], &yy[i], &zz[i])
	}
	cnt1 := 0
	var p, q int
	for k := 0; k < n; k++ {
		for l := k + 1; l < n; l++ {
			if xx[k] == xx[l] {
				continue
			}
			var bad [N]bool
			cnt := 0
			for i := 0; i < n; i++ {
				if i == k || i == l || collinear(i, k, l) {
					bad[i] = true
					cnt++
				}
			}
			cnt--
			m := 0
			for i := 0; i < n; i++ {
				if !bad[i] {
					for j := i + 1; j < n; j++ {
						if !bad[j] && coplanar(i, j, k, l) {
							x1 := xx[j] - xx[i]
							y1 := yy[j] - yy[i]
							z1 := zz[j] - zz[i]
							x2 := xx[l] - xx[k]
							y2 := yy[l] - yy[k]
							z2 := zz[l] - zz[k]
							if q = cross2(x1, y1, x2, y2); q != 0 {
								p = cross2(x1, y1, xx[i]-xx[k], yy[i]-yy[k])
							} else if q = cross2(y1, z1, y2, z2); q != 0 {
								p = cross2(y1, z1, yy[i]-yy[k], zz[i]-zz[k])
							} else if q = cross2(z1, x1, z2, x2); q != 0 {
								p = cross2(z1, x1, zz[i]-zz[k], xx[i]-xx[k])
							} else {
								continue
							}
							if q < 0 {
								p = -p
								q = -q
							}
							d := gcd(abs(p), q)
							p /= d
							q /= d
							if xx[k]*q+x2*p >= 0 {
								continue
							}
							pp[m] = p
							qq[m] = q
							ii[m] = i
							jj[m] = j
							m++
						}
					}
				}
			}
			for h := 0; h < m; h++ {
				hh[h] = h
			}
			sort(hh, 0, m)
			cnt1 = max(cnt1, cnt)
			var h_, cnt_ int
			for h := 0; h < m; h = h_ {
				h_ = h
				cnt_ = cnt
				for h_ < m && pp[hh[h_]] == pp[hh[h]] && qq[hh[h_]] == qq[hh[h]] {
					if h_ == h || ii[hh[h_]] != ii[hh[h_-1]] {
						cnt_++
					}
					h_++
				}
				cnt1 = max(cnt1, cnt_)
			}
		}
	}
	fmt.Println(n - cnt1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
