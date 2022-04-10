package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

type node struct {
	w, h  int
	isBox int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	A := make([]int, N)
	B := make([]int, N)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	for i := range B {
		fmt.Fscan(in, &B[i])
	}
	all := []node{}
	allLen := make(map[int]bool)
	for i := 0; i < N; i++ {
		all = append(all, node{A[i], B[i], 1})
		allLen[B[i]] = true
	}

	C := make([]int, M)
	D := make([]int, M)
	for i := range C {
		fmt.Fscan(in, &C[i])
	}
	for i := range D {
		fmt.Fscan(in, &D[i])
	}
	for i := 0; i < M; i++ {
		all = append(all, node{C[i], D[i], 0})
		allLen[D[i]] = true
	}

	allLength := []int{}
	for k := range allLen {
		allLength = append(allLength, k)
	}
	sort.Ints(allLength)

	convert := make(map[int]int)
	for i := range allLength {
		convert[allLength[i]] = i
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].w == all[j].w {
			return all[i].isBox < all[j].isBox
		}
		return all[i].w > all[j].w
	})

	bit := NewBIT(N + M)
	for i := range all {
		if all[i].isBox == 0 {
			bit.Add(convert[all[i].h], 1)
		} else {
			j := bit.maxMinIndex(convert[all[i].h])
			if j == len(all) {
				fmt.Println("No")
				return
			} else {
				bit.Add(j, -1)
			}
		}
	}
	fmt.Println("Yes")
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

func (b BIT) UpperBound(x int) int {
	idx, y := 0, 0
	n := bits.Len(uint(len(b) - 1))
	for i := n; i > -1; i-- {
		k := idx + (1 << i)
		if k <= len(b)-1 && (y+b[k] <= x) {
			y += b[k]
			idx += 1 << i
		}
	}
	return idx
}

func (b BIT) querry(i int) (ans int) {
	for i > 0 {
		ans += b[i]
		i -= -i & i
	}
	return
}

func (b BIT) maxMinIndex(i int) int {
	return b.UpperBound(b.querry(i))
}
