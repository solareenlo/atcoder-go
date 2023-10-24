package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	type pair struct {
		x, y int
	}

	var num, query int
	fmt.Fscan(in, &num, &query)
	pls := make([]int, query)
	idx := make([]int, query)
	val := make([]int, query)
	maxaa := 0
	for i := 0; i < query; i++ {
		fmt.Fscan(in, &pls[i], &idx[i], &val[i])
		idx[i]--
		val[i]--
		maxaa = max(maxaa, pls[i])
	}
	var com [111][111]int
	for i := 0; i < 111; i++ {
		com[i][0] = 1
		for j := 1; j <= i; j++ {
			com[i][j] = (com[i-1][j-1] + com[i-1][j]) % MOD
		}
	}
	if num <= 6 {
		v := make([][]int, 0)
		z := make([]int, 0)
		for i := 0; i < num; i++ {
			z = append(z, i)
		}
		for {
			tmp := make([]int, len(z))
			copy(tmp, z)
			v = append(v, tmp)
			if !nextPermutation(sort.IntSlice(z)) {
				break
			}
		}
		cnt := 0
		for i := 0; i < len(v); i++ {
			f := true
			for j := 0; j < query; j++ {
				if i+pls[j] >= len(v) || v[i+pls[j]][idx[j]] != val[j] {
					f = false
				}
			}
			if f {
				cnt++
			}
		}
		fmt.Println(cnt)
		return
	}
	ans := 0
	for l := 6; l <= num; l++ {
		for p := 0; p < l; p++ {
			pos := make([]int, l)
			pos[0] = p
			pt := 1
			for i := l - 1; i >= 0; i-- {
				if p == i {
					continue
				}
				pos[pt] = i
				pt++
			}
			reverseOrderInt(pos[l-5 : l])
			v := make([][]int, 0)
			for i := 0; i < 180; i++ {
				tmp := make([]int, len(pos))
				copy(tmp, pos)
				v = append(v, tmp)
				if !nextPermutation(sort.IntSlice(pos)) {
					break
				}
			}
			if p == l-1 && l != num {
				continue
			}
			cnt := 0
			for st := 0; st < 120; st++ {
				f := true
				d := make([]int, 60)
				for i := range d {
					d[i] = -1
				}
				for i := 0; i < query; i++ {
					if idx[i] < num-l {
						continue
					}
					if st+pls[i] >= len(v) {
						f = false
						break
					}
					x := v[st+pls[i]][idx[i]-(num-l)]
					if d[x] != -1 && d[x] != val[i] {
						f = false
						break
					}
					d[x] = val[i]
				}
				if !f {
					continue
				}
				isval := make([]int, 60)
				for i := range isval {
					isval[i] = 1
				}
				for i := 0; i < query; i++ {
					if idx[i] < num-l {
						isval[val[i]] = 0
					}
				}
				for i := 0; i < l; i++ {
					if d[i] != -1 && isval[d[i]] == 0 {
						f = false
					}
				}
				if !f {
					continue
				}
				ls := make([]pair, 0)
				ls = append(ls, pair{-1, -1})
				for i := 0; i < l; i++ {
					if d[i] != -1 {
						ls = append(ls, pair{i, d[i]})
					}
				}
				ls = append(ls, pair{l, num})
				rr := 1
				for i := 1; i < len(ls); i++ {
					sum := 0
					req := ls[i].x - ls[i-1].x - 1
					for j := ls[i-1].y + 1; j < ls[i].y; j++ {
						sum += isval[j]
					}
					if ls[i-1].y >= ls[i].y {
						rr = 0
					}
					if sum < req {
						rr = 0
					} else {
						rr = rr * com[sum][req] % MOD
					}
				}
				cnt += rr
			}
			d := make([]int, 60)
			for i := range d {
				d[i] = -1
			}
			ff := true
			for i := 0; i < query; i++ {
				if idx[i] < num-l {
					if d[idx[i]] != -1 && d[idx[i]] != val[i] {
						ff = false
					}
					d[idx[i]] = val[i]
				}
			}
			if ff {
				c := 0
				for i := 0; i < num-l; i++ {
					if d[i] == -1 {
						c++
					}
				}
				for i := 1; i <= c; i++ {
					cnt = cnt * i % MOD
				}
				ans = (ans + cnt) % MOD
			}
		}
	}
	fmt.Println(ans % MOD)
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
