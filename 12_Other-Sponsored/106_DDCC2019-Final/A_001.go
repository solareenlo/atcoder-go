package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	ma, cnt := 0.0, 0.0
	ans := 0.0
	for i := 0; i < n; i++ {
		if s[i] == '>' {
			cnt++
			ma = math.Max(ma, cnt)
		} else {
			cnt = 0
		}
		ans += 1 / (cnt + 1)
	}
	ans -= (ma + 1) / (ma + 2)
	fmt.Println(ans)
}
