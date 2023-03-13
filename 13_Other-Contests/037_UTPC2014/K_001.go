package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solve())
}

func solve() int {
	const INF = int(1e18)

	var a, b, x int
	fmt.Scan(&a, &b, &x)
	if a == x {
		return 0
	}

	L := 36
	H := 1 << 18
	res := 0
	for L > 0 && (b>>(L-1)&1) == 0 {
		if a == x {
			return res
		}
		if (x >> (L - 1) & 1) != 0 {
			return -1
		}
		if ((a >> (L - 1)) & 1) != 0 {
			a = (a / 2) ^ ((a & 1) * b)
			res++
		}
		L--
	}
	H = 1 << ((L + 1) / 2)
	if a == x {
		return res
	}

	var conv_H [36]int

	for i := 0; i < L; i++ {
		cur := 1 << i
		for j := 0; j < H; j++ {
			cur = (cur / 2) ^ ((cur & 1) * b)
		}
		conv_H[i] = cur
	}

	var baby [1 << 18]pair
	cur := x
	for i := 0; i < H; i++ {
		baby[i] = pair{cur, i}
		cur = (cur / 2) ^ ((cur & 1) * b)
	}
	tmp := baby[:H]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})

	ans := INF

	cur = a
	for i := 0; i < H; i++ {
		c := 0
		for j := 0; j < L; j++ {
			if ((cur >> j) & 1) != 0 {
				c ^= conv_H[j]
			}
		}

		tmp := baby[:H]
		pos := lowerBound(tmp, pair{c + 1, -1})
		pos--
		if pos >= 0 && baby[pos].x == c {
			ans = min(ans, H*(i+1)-baby[pos].y)
		}

		cur = c
	}
	if ans == INF {
		return -1
	}
	return ans + res
}

type pair struct {
	x, y int
}

func lowerBound(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x
	})
	return idx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
