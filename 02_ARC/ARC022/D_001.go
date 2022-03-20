package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"sort"
)

type pair struct{ x, y int }

func veccross(p1, p2, p3 pair) int {
	p3.x -= p1.x
	p2.x -= p1.x
	p3.y -= p1.y
	p2.y -= p1.y
	return p3.x*p2.y - p2.x*p3.y
}

func resize(a []int, n int) []int {
	res := make([]int, n)
	if len(a) < n {
		copy(res, a)
	} else {
		res = a[:n]
	}
	return res
}

func convex_hull(vp []pair) []int {
	type pair2 struct {
		p pair
		i int
	}
	sorted := make([]pair2, 0)
	res := make([]int, 0)

	if len(vp) <= 2 {
		if len(vp) >= 1 {
			res = append(res, 0)
		}
		if len(vp) >= 2 {
			res = append(res, 1)
		}
		return res
	}

	for i := 0; i < len(vp); i++ {
		sorted = append(sorted, pair2{vp[i], i})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].p.x < sorted[j].p.x || (sorted[i].p.x == sorted[j].p.x && sorted[i].p.y < sorted[j].p.y)
	})
	res = resize(res, len(vp)*2)
	k := 0
	for i := 0; i < len(vp); i++ {
		for k > 1 && veccross(vp[res[k-2]], vp[res[k-1]], sorted[i].p) <= 0 {
			k--
		}
		res[k] = sorted[i].i
		k++
	}
	rb := k
	for i := len(vp) - 2; i >= 0; i-- {
		for k > rb && veccross(vp[res[k-2]], vp[res[k-1]], sorted[i].p) <= 0 {
			k--
		}
		res[k] = sorted[i].i
		k++
	}
	res = resize(res, k-1)
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	V := make([]pair, 0)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		V = append(V, pair{x, y})
	}

	C := convex_hull(V)
	CC := make([]complex128, 0)

	for i := 0; i < len(C); i++ {
		cx1 := V[C[i]].x
		cy1 := V[C[i]].y
		cx2 := V[C[(i+1)%len(C)]].x
		cy2 := V[C[(i+1)%len(C)]].y
		c1 := complex(float64(cx1), float64(cy1))
		c2 := complex(float64(cx2), float64(cy2))

		v1 := c2 - c1
		v2 := -c1
		var v3 complex128
		ab := cmplx.Abs(v1)
		v1 = complex(real(v1)/ab, imag(v1)/ab)
		dot := real(v1)*real(v2) + imag(v1)*imag(v2)
		v3 = v2 - complex(dot, 0)*v1
		if real(v1)*imag(v2)-real(v2)*imag(v1) > 0 {
			v3 = -v3
		}
		v2 = complex(dot, 0)*v1 - v3 + c1
		CC = append(CC, v2)
	}

	cond := make([][]pair, 500002)
	for i := 0; i < len(C); i++ {
		cx1 := V[C[i]].x
		cy1 := V[C[i]].y
		c := complex(float64(cx1), float64(cy1))
		p1 := CC[(i+len(C)-1)%len(C)] - c
		p2 := CC[i] - c

		r := cx1*cx1 + cy1*cy1
		if p1 == p2 {
			for y := -int(math.Sqrt(float64(r))) - 1; float64(y) <= (math.Sqrt(float64(r)))+1; y++ {
				if int(r)-y*y >= 0 {
					x := int(math.Sqrt(float64(r - y*y)))
					if (x+1)*(x+1)+y*y <= r {
						x++
					}
					cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 - x, 1000000 + i})
					cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 + x + 1, i})
				}
			}
		} else if real(p1) <= 0 && real(p2) <= 0 {
			if imag(p1) <= imag(p2) {
				for y := int(math.Ceil(imag(p1))); float64(y) <= imag(p2); y++ {
					if int(r)-y*y >= 0 {
						x := int(math.Sqrt(float64(r - y*y)))
						if (x+1)*(x+1)+y*y <= r {
							x++
						}
						cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 - x, 1000000 + i})
						if x == 0 {
							cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 + x + 1, 0 + i})
						}
					}
				}
			} else {
				for y := -int(math.Sqrt(float64(r))) - 1; y <= int(math.Sqrt(float64(r)))+1; y++ {
					if int(r)-y*y >= 0 {
						x := int(math.Sqrt(float64(r - y*y)))
						if (x+1)*(x+1)+y*y <= r {
							x++
						}
						cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 + x + 1, i})
						if y >= int(imag(p2)) || y <= int(imag(p1)) {
							cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 - x, 1000000 + i})
						}
					}
				}
			}
		} else if real(p1) > 0 && real(p2) < 0 {
			for y := int(math.Floor(imag(p1))); y >= -int(math.Sqrt(float64(r)))-1; y-- {
				if r-y*y >= 0 {
					x := int(math.Sqrt(float64(r - y*y)))
					if (x+1)*(x+1)+y*y <= r {
						x++
					}
					cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 + x + 1, i})
				}
			}
			for y := -int(math.Sqrt(float64(r))) - 1; float64(y) <= imag(p2); y++ {
				if r-y*y >= 0 {
					x := int(math.Sqrt(float64(r - y*y)))
					if (x+1)*(x+1)+y*y <= r {
						x++
					}
					cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 - x, 1000000 + i})
				}
			}
		} else if real(p1) < 0 && real(p2) > 0 {
			for y := int(math.Ceil(imag(p1))); y <= int(math.Sqrt(float64(r))+1); y++ {
				if r-y*y >= 0 {
					x := int(math.Sqrt(float64(r - y*y)))
					if (x+1)*(x+1)+y*y <= r {
						x++
					}
					cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 - x, 1000000 + i})
				}
			}
			for y := int(math.Sqrt(float64(r))) + 1; float64(y) >= imag(p2); y-- {
				if r-y*y >= 0 {
					x := int(math.Sqrt(float64(r - y*y)))
					if (x+1)*(x+1)+y*y <= r {
						x++
					}
					cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 + x + 1, i})
				}
			}
		} else if real(p1) > 0 && real(p2) > 0 {
			if imag(p1) > imag(p2) {
				for y := int(math.Floor(imag(p1))); float64(y) >= imag(p2); y-- {
					if r-y*y >= 0 {
						x := int(math.Sqrt(float64(r - y*y)))
						if (x+1)*(x+1)+y*y <= r {
							x++
						}
						cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 + x + 1, i})
					}
				}
			} else {
				for y := -int(math.Sqrt(float64(r))) - 1; y <= int(math.Sqrt(float64(r)))+1; y++ {
					if r-y*y >= 0 {
						x := int(math.Sqrt(float64(r - y*y)))
						if (x+1)*(x+1)+y*y <= r {
							x++
						}
						cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 - x, 1000000 + i})
						if y <= int(imag(p2)) || y >= int(imag(p1)) {
							cond[250000+cy1+y] = append(cond[250000+cy1+y], pair{cx1 + x + 1, i})
						}
					}
				}
			}
		}
	}

	ret := 0
	for i := 0; i < 500001; i++ {
		if len(cond[i]) > 0 {
			sort.Slice(cond[i], func(j, k int) bool {
				return cond[i][j].x < cond[i][k].x || (cond[i][j].x == cond[i][k].x && cond[i][j].x < cond[i][k].x)
			})
			pre := -1000000
			st := 0
			for j := 0; j < len(cond[i]); j++ {
				if j > 0 && (cond[i][j].x == cond[i][j-1].x) && (cond[i][j].y/1000000 == cond[i][j-1].y/1000000) {
					continue
				}
				if cond[i][j].y >= 1000000 {
					if st == 0 {
						pre = cond[i][j].x
					}
					st++
				} else {
					st--
					if st == 0 {
						ret += cond[i][j].x - pre
					}
				}
			}
		}
	}
	fmt.Println(ret)
}
