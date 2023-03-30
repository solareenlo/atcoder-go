package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	fact := make([]int, n+1)
	fact[0] = 1
	fw := NewInt(n + 1)
	res := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		fact[i+1] = (i + 1) * fact[i] % mod
		fw.Add(i+1, 1)
	}
	for i := 0; i < n; i++ {
		res += fact[n-1-i] * fw.Sum(0, a[i]) % mod
		res %= mod
		fw.Add(a[i], -1)
	}
	fmt.Println(res + 1)
}

// Ref: https://github.com/monkukui/ac-library-go
type Int struct {
	n    int
	data []uint
}

func NewInt(n int) *Int {
	i := &Int{
		n:    n,
		data: make([]uint, n),
	}
	for idx := range i.data {
		i.data[idx] = 0
	}
	return i
}

func (i *Int) Add(pos int, x int) {
	if !(0 <= pos && pos < i.n) {
		panic("")
	}
	pos++
	for pos <= i.n {
		i.data[pos-1] += uint(x)
		pos += pos & -pos
	}
}

func (i *Int) Sum(l, r int) int {
	if !(0 <= l && l <= r && r <= i.n) {
		panic("")
	}
	return int(i.sum(r) - i.sum(l))
}

func (i *Int) sum(r int) uint {
	s := uint(0)
	for r > 0 {
		s += i.data[r-1]
		r -= r & -r
	}
	return s
}
