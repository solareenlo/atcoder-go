package main

import (
	"bufio"
	"fmt"
	"os"
)

func addFenwick(beg []int, n uint, i uint, val int) {
	for ; i < n; i |= i + 1 {
		beg[i] += val
	}
}

func sumFenwick(beg []int, i int) int {
	sum := 0
	for ; i > 0; i = i & (i - 1) {
		sum += beg[i-1]
	}
	return sum
}

func highestOneBit(v uint) uint {
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	return (v >> 1) + 1
}

func lowerBoundFenwick(beg []int, n int, val int) int {
	left := 0
	if !(left < val) {
		return 0
	}
	var i int = 0
	for w := int(highestOneBit(uint(n))); w > 0; w >>= 1 {
		if i+w <= n {
			mid := left
			mid += beg[i+w-1]
			if mid < val {
				i += w
				left = mid
			}
		}
	}
	return i + 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, Q int
	fmt.Fscan(reader, &N, &Q)
	ft := make([]int, N-1)
	base := 0
	for i := 0; i < Q; i++ {
		var ty, p, v int
		fmt.Fscan(reader, &ty)
		if ty == 1 {
			fmt.Fscan(reader, &p, &v)
			m := i + 1 + v
			t := i + 1
			for d := base / m; ; d++ {
				u := d*m + t - base
				j := lowerBoundFenwick(ft, len(ft), u)
				if j == N {
					break
				}
				for k := max(N-p, j); k < N; {
					val := base + sumFenwick(ft, k)
					if val >= ((d + 1) * m) {
						break
					}
					newval := (val + m) / m * m
					diff := newval - val
					l := lowerBoundFenwick(ft, len(ft), val-base+1)
					if k == 0 {
						base += diff
					} else {
						addFenwick(ft, uint(len(ft)), uint(k-1), diff)
					}
					if l < N {
						addFenwick(ft, uint(len(ft)), uint(l-1), -diff)
					}
					k = l
				}
			}
		} else if ty == 2 {
			fmt.Fscan(reader, &p)
			p--
			ans := base + sumFenwick(ft, N-1-p)
			fmt.Println(ans)
		} else if ty == 3 {
			fmt.Fscan(reader, &v)
			base += v
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
