package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type P struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	L := make([]int, n)
	R := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &L[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &R[i])
	}

	var comp func(P, P, P) bool
	comp = func(p, q, r P) bool {
		return (q.y-p.y)*(r.x-q.x) > (r.y-q.y)*(q.x-p.x)
	}

	id := make([]P, 0)
	for i := 0; i < n; i++ {
		p := P{i, R[i]}
		for len(id) >= 2 && comp(id[len(id)-2], id[len(id)-1], p) {
			id = id[:len(id)-1]
		}
		id = append(id, p)
	}
	pos := 0
	for i := 0; i < n; i++ {
		if pos+1 < len(id) && id[pos+1].x <= i {
			pos++
		}
		if id[pos].x == i {
			continue
		}
		p := id[pos]
		q := id[pos+1]
		r := P{i, L[i]}
		if comp(p, r, q) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
