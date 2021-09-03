package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a int
	fmt.Scan(&n, &a)
	var k string
	fmt.Scan(&k)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&b[i])
	}
	step := make([]int, n+1)
	for i := range step {
		step[i] = -1
	}

	now, cnt := a, 0
	for {
		if step[now] != -1 {
			break
		}
		step[now] = cnt
		cnt++
		now = b[now]
	}

	kModc, t := 0, step[now]
	cycle := cnt - t
	for i := 0; i < len(k); i++ {
		kModc = (kModc*10 + int(k[i]-'0')) % cycle
	}
	for kModc < t {
		kModc += cycle
	}
	i, _ := strconv.Atoi(k)
	if len(k) < 9 && i <= n {
		kModc = i
	}

	res := a
	for ; kModc > 0; kModc-- {
		res = b[res]
	}
	fmt.Println(res)
}
