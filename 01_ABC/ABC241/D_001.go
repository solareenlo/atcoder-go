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

	var q int
	fmt.Fscan(in, &q)
	val := make([]int, q)
	a := make([][3]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &a[i][0], &a[i][1])
		if a[i][0] != 1 {
			fmt.Fscan(in, &a[i][2])
		}
		val[i] = a[i][1]
	}

	val = append(val, 0)

	id := make(map[int]int)
	sort.Ints(val)
	for i, v := range val {
		id[v] = i
	}

	bit := NewBIT(q + 1)
	for i := 0; i < q; i++ {
		if a[i][0] == 1 {
			bit.Add(id[a[i][1]], 1)
		} else if a[i][0] == 2 {
			s := bit.Sum(id[a[i][1]])
			id := bit.LowerBound(s - a[i][2] + 1)
			if val[id] == 0 {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, val[id])
			}
		} else {
			s := bit.Sum(id[a[i][1]] - 1)
			id := bit.LowerBound(s + a[i][2])
			if id == q+1 {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, val[id])
			}
		}
	}
}

const mod = 998244353

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

func (t BIT) LowerBound(x int) int {
	idx, k := 0, 1
	for k < len(t) {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k < len(t) && t[idx+k] < x {
			x -= t[idx+k]
			idx += k
		}
	}
	return idx
}
