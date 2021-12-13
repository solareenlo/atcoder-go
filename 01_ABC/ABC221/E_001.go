package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func compress(A []int) int {
	m := map[int]int{}
	for i := range A {
		m[A[i]] = 0
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	size := 0
	for i := range keys {
		m[keys[i]] = size
		size++
	}
	for i := range A {
		A[i] = m[A[i]]
	}
	return size
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}
	n := compress(a)
	bit := NewBIT(n)

	div := powMod(2, mod-2)
	res := 0
	for i := 0; i < N; i++ {
		res += bit.Sum(a[i]) * powMod(2, i) % mod
		res %= mod
		bit.Add(a[i], powMod(div, i+1))
	}
	fmt.Println(res)
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

type BIT []int

func (BIT) e() int {
	return 0
}

func NewBIT(n int) BIT {
	bit := make(BIT, n+1)
	for i := range bit {
		bit[i] = bit.e()
	}
	return bit
}

func (BIT) op(a, b int) int {
	return (a + b) % mod
}

func (b BIT) Add(i, v int) {
	i++
	for ; i < len(b); i += i & -i {
		b[i] = b.op(b[i], v)
	}
}

func (b BIT) Sum(i int) int {
	i++
	sum := b.e()
	for ; 0 < i; i -= i & -i {
		sum = b.op(sum, b[i])
	}
	return sum
}
