package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const inf = 1001001001001001001

	type pair struct {
		x, y int
	}

	var n, E int
	fmt.Fscan(in, &n, &E)
	v := make([]int, n)
	rev := make([]int, n*2)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	idx := make([]pair, n*2)
	sp := make([]bool, n)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		r--
		idx[l] = pair{i, 0}
		idx[r] = pair{i, 1}
		rev[r] = l
	}

	lst := -1
	for i := 0; i < n*2; i++ {
		if idx[i].y == 1 {
			if lst < rev[i] {
				lst = rev[i]
				sp[idx[i].x] = true
			}
		}
	}

	ok, ng := -1, 1000000000001
	for ng-ok > 1 {
		md := (ok + ng) / 2
		e := E
		bank := make([]int, 0)
		cur := make([]int, len(v))
		copy(cur, v)
		able := true
		for i := 0; i < n*2; i++ {
			if e > inf {
				break
			}
			if idx[i].y == 0 {
				if sp[idx[i].x] {
					cur[idx[i].x] += e
					e = 0
					bank = append(bank, idx[i].x)
				} else {
					e += cur[idx[i].x]
					cur[idx[i].x] = 0
				}
			} else {
				if sp[idx[i].x] {
					cur[idx[i].x] *= 2
				}
				cur[idx[i].x] += e
				e = 0
				for len(bank) != 0 && cur[idx[i].x] < md {
					if bank[len(bank)-1] == idx[i].x {
						bank = bank[:len(bank)-1]
						continue
					}
					if cur[idx[i].x]+cur[bank[len(bank)-1]] > md {
						cur[bank[len(bank)-1]] -= md - cur[idx[i].x]
						cur[idx[i].x] = md
					} else {
						cur[idx[i].x] += cur[bank[len(bank)-1]]
						cur[bank[len(bank)-1]] = 0
						bank = bank[:len(bank)-1]
					}
				}
				if cur[idx[i].x] < md {
					able = false
					break
				}
				e = cur[idx[i].x] - md
				cur[idx[i].x] = 0
			}
		}
		if able {
			ok = md
		} else {
			ng = md
		}
	}
	if ok == 1000000000000 {
		fmt.Println("inf")
	} else {
		fmt.Println(ok)
	}
}
