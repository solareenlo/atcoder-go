package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	type hzy struct{ x, id int }
	a := make([]hzy, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].x)
		a[i].id = i
	}
	a = a[1:]
	sort.Slice(a, func(i, j int) bool {
		return a[i].x > a[j].x
	})
	tmp := make([]hzy, 1)
	a = append(tmp, a...)

	ans := 0
	b := make([]int, 200005)
	for i := 1; i <= n/2; i++ {
		ans += a[i].x
		b[a[i].id] = -1
	}

	s := 0
	min1 := 0
	ans1 := 0
	for i := 1; i <= n; i++ {
		s += b[i]
		if b[i] == 0 {
			s++
		}
		if s < min1 {
			min1 = s
			ans1 = i
		}
	}
	fmt.Println(ans1, ans)
}
