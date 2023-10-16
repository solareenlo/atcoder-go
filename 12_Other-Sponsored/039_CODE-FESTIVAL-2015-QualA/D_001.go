package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int64
	fmt.Fscan(in, &n, &m)

	x := make([]int64, m)
	for i := 0; i < int(m); i++ {
		fmt.Fscan(in, &x[i])
	}

	if m == n {
		fmt.Println("0")
		return
	}

	r := int64(0)
	l := 2 * n

	for l-r > 1 {
		s := (l + r) / 2
		cur := int64(0)
		flag := true

		for i := 0; i < int(m); i++ {
			if x[i]-cur-1 > s {
				continue
			} else if cur+1 >= x[i] {
				cur = max(cur, x[i]+s)
			} else {
				cur = max((s-(x[i]-cur-1))/2+x[i], s-(x[i]-cur-1)*2+x[i])
			}
		}

		if flag && cur < n {
			flag = false
		}

		if flag {
			l = s
		} else {
			r = s
		}
	}

	fmt.Println(l)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
