package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 314159265

type pair struct {
	x, y int
}

var p []pair
var S [20][(1 << 17) + 5]int

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100005

	var n, d int
	fmt.Fscan(in, &n, &d)

	p = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})

	for i := 0; i < n; i++ {
		p[i].x = p[i].y
		p[i].y = n - i
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})

	cha(0, 1, 1)
	for i := 0; i < n; i++ {
		for j := d - 1; j >= 0; j-- {
			cha(j+1, p[i].y, ret(j, p[i].y))
		}
	}
	fmt.Println(S[d][1<<17])
}

func cha(a, b, c int) {
	for b <= 1<<17 {
		S[a][b] = (S[a][b] + c) % mod
		b += (b & -b)
	}
}

func ret(a, b int) int {
	num := 0
	for b != 0 {
		num = (num + S[a][b]) % mod
		b -= b & -b
	}
	return num
}
