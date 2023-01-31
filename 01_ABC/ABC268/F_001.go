package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1000005

	var n int
	fmt.Fscan(in, &n)

	a := make([]node, N)
	ans := 0
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = " " + s
		l := len(s)
		for j := 1; j < l; j++ {
			if s[j] == 'X' {
				a[i].x++
			} else {
				a[i].y += int(s[j] - '0')
				ans += int(s[j]-'0') * a[i].x
			}
		}
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].x*tmp[j].y > tmp[j].x*tmp[i].y
	})

	p := 0
	for i := 1; i <= n; i++ {
		ans += p * a[i].y
		p += a[i].x
	}
	fmt.Println(ans)
}
