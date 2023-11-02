package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func above(p, q, r P) bool {
	if p.a == r.a {
		return true
	}
	return (q.a-p.a)*(r.b-p.b) <= (q.b-p.b)*(r.a-p.a)
}

func below(p, q, r P) bool {
	if p.a == r.a {
		return true
	}
	return (q.a-p.a)*(r.b-p.b) >= (q.b-p.b)*(r.a-p.a)
}

func dist(p, q P) float64 {
	return math.Sqrt((p.a-q.a)*(p.a-q.a) + (p.b-q.b)*(p.b-q.b))
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var st [200005]P

	var n int
	fmt.Fscan(in, &n)
	vec := make([]P, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &vec[i].a, &vec[i].b)
	}
	var s, t P
	fmt.Fscan(in, &s.a, &s.b)
	fmt.Fscan(in, &t.a, &t.b)
	all := make([]P, n)
	for i := range all {
		all[i].a = vec[i].a
		all[i].b = vec[i].b
	}
	all = append(all, s)
	all = append(all, t)
	sortPair(all)
	sz := 0
	for i := 0; i < len(all); i++ {
		for sz >= 2 && above(st[sz-2], st[sz-1], all[i]) {
			sz--
		}
		st[sz] = all[i]
		sz++
	}
	ls := sz
	for i := len(all) - 2; i >= 0; i-- {
		for sz >= ls+1 && below(all[i], st[sz-1], st[sz-2]) {
			sz--
		}
		st[sz] = all[i]
		sz++
	}
	sz--
	c := 0
	a := -1
	b := -1
	for i := 0; i < sz; i++ {
		if st[i] == s || st[i] == t {
			c++
			if a == -1 {
				a = i
			} else {
				b = i
			}
		}
	}
	if c == 1 {
		fmt.Println(dist(s, t))
	} else {
		sum := 0.0
		sum2 := 0.0
		for i := 0; i < sz; i++ {
			d := dist(st[i], st[(i+1)%sz])
			sum += d
			if a <= i && i < b {
				sum2 += d
			}
		}
		fmt.Println(math.Min(sum2, sum-sum2))
	}
}

type P struct {
	a, b float64
}

func sortPair(tmp []P) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].a == tmp[j].a {
			return tmp[i].b < tmp[j].b
		}
		return tmp[i].a < tmp[j].a
	})
}
