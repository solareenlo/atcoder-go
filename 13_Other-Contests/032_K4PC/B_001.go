package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAX = 100005

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	c := make([]int, MAX)
	v := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
		v = append(v, c[i])
	}
	v = append(v, m)
	sort.Ints(v)
	v = unique(v)
	t := make([]int, MAX)
	for i := 0; i < len(v); i++ {
		t[i+1] = v[i]
	}
	m = lowerBound(v, m)
	m++
	bit := NewBit()
	bit.add(m)
	bit1 := NewBit()
	bit1.add(len(v) + 2 - m)

	cost := 0
	for i := 0; i < n; i++ {
		p := lowerBound(v, c[i])
		p++
		k := bit.sum(p)
		k1 := bit1.sum(len(v) + 2 - p)
		if k1 != -1 {
			k1 = len(v) + 2 - k1
		}
		kk := 0
		if k1 == -1 {
			kk = k
		} else {
			if k == -1 {
				kk = k1
			} else {
				del := abs(t[k] - t[p])
				dell := abs(t[k1] - t[p])
				if del > dell {
					kk = k1
				} else {
					kk = k
				}
			}
		}
		cost += abs(t[p] - t[kk])
		bit.add(p)
		bit1.add(len(v) + 2 - p)
	}
	fmt.Println(cost)
}

type Bit struct {
	bit []int
}

func NewBit() *Bit {
	b := new(Bit)
	b.bit = make([]int, MAX)
	for i := range b.bit {
		b.bit[i] = -1
	}
	return b
}

func (b *Bit) add(i int) {
	x := i
	for i < MAX {
		b.bit[i] = max(b.bit[i], x)
		i += i & -i
	}
}

func (b Bit) sum(i int) int {
	res := -1
	for i > 0 {
		res = max(res, b.bit[i])
		i -= i & -i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
