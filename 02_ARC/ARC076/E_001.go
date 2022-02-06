package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	R int
	C int
)

const INF = 1 << 60

func id(x, y int) int {
	if x == 0 {
		return y
	}
	if x == R {
		return -R - y
	}
	if y == 0 {
		return -x
	}
	if y == C {
		return C + x
	}
	return INF
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &R, &C, &n)

	type pair struct{ x, y int }
	a := make([]pair, 200200)
	cnt := 0
	for i := 1; i <= n; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		x := id(x1, y1)
		y := id(x2, y2)
		if x != INF && y != INF {
			cnt++
			a[cnt] = pair{x, i}
			cnt++
			a[cnt] = pair{y, i}
		}
	}
	b := a[1 : cnt+1]
	sort.Slice(b, func(i, j int) bool {
		return b[i].x < b[j].x || (b[i].x == b[j].x && b[i].y < b[j].y)
	})

	t := 0
	s := make([]pair, 200200)
	for i := 1; i <= cnt; i++ {
		if t == 0 || s[t].y != a[i].y {
			t++
			s[t] = a[i]
		} else {
			t--
		}
	}
	if t == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
