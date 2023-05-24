package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var p, q []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	p = make([]int, n+2)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &p[i])
	}
	q = make([]int, n+2)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &q[i])
	}
	sp := make([]int, 0)
	sq := make([]int, 0)
	solve(p, &sp)
	solve(q, &sq)
	for i := 1; i < n+1; i++ {
		if p[i] != q[i] {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(len(sq) + len(sp))
	sq = reverseOrderInt(sq)
	for i := range sp {
		fmt.Printf("%d ", sp[i])
	}
	for i := range sq {
		fmt.Printf("%d ", sq[i])
	}
	fmt.Println()
}

func solve(p []int, ans *[]int) {
	y := make([][]int, 0)
	for i := 1; i < n+1; i++ {
		cur := 0
		ptr := i
		for len(y) > cur {
			if y[cur][len(y[cur])-1] < p[ptr] {
				y[cur] = append(y[cur], p[ptr])
				break
			}
			sz := len(y[cur])
			s := sz - 1
			for s > 0 && y[cur][s-1] > p[ptr] {
				s--
			}
			for i := sz - 1; i > s; i-- {
				ptr--
				*ans = append(*ans, ptr)
				p[ptr], p[ptr+1] = p[ptr+1], p[ptr]
			}
			y[cur][s] = p[ptr]
			ptr--
			for i := s - 1; i >= 0; i-- {
				ptr--
				*ans = append(*ans, ptr)
				p[ptr], p[ptr+1] = p[ptr+1], p[ptr]
			}
			cur++
		}
		if len(y) == cur {
			v := make([]int, 0)
			v = append(v, p[ptr])
			y = append(y, v)
		}
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
