package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	t := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		t = append(t, a[i])
	}

	var q int
	fmt.Fscan(in, &q)
	type P struct{ l, r, x int }
	p := make([]P, q)
	for i := 0; i < q; i++ {
		var l, r, x int
		fmt.Fscan(in, &l, &r, &x)
		l--
		p[i] = P{l, r, x}
		t = append(t, x)
	}
	sort.Ints(t)
	unique(t)
	t = erase(t, len(t)-1)
	m := len(t)

	for i := 0; i < n; i++ {
		a[i] = lowerBound(t, a[i])
	}
	for i := 0; i < q; i++ {
		p[i].x = lowerBound(t, p[i].x)
	}

	c := make([]int, m+1)
	s := 0
	for i := 0; i < n; i++ {
		c[a[i]]++
	}
	for j := 0; j < m; j++ {
		s += c[j] * (c[j] - 1) / 2
	}

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[i] = a[i]
	}
	mp[n] = 0

	for i := range p {
		l := p[i].l
		r := p[i].r
		x := p[i].x
		keys := make([]int, 0, len(mp))
		for k := range mp {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		i0 := upperBound(keys, l)
		i0--
		i1 := lowerBound(keys, r)
		for keys[i0] <= keys[i1] && i0 < i1 {
			i := keys[i0]
			a := mp[keys[i0]]
			delete(mp, keys[i0])
			i0++
			j := keys[i0]
			c1 := c[a]
			c[a] -= j - i
			if i < l {
				mp[i] = a
				c[a] += l - i
			}
			if r < j {
				mp[r] = a
				c[a] += j - r
			}
			s -= (c1 - c[a]) * (c[a] + c1 - 1) / 2
		}
		mp[l] = x
		c0 := c[x]
		c[x] += r - l
		s += (r - l) * (c0 + c[x] - 1) / 2
		fmt.Fprintln(out, s)
	}
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func erase(a []int, pos int) []int {
	return append(a[:pos], a[pos+1:]...)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
