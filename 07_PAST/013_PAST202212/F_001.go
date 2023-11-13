package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b, c, d int
	var s string
	fmt.Scan(&n, &a, &b, &c, &d, &s)
	x, _ := strconv.Atoi(string(s[0]))
	for i := 2; i < 5; i++ {
		tmp, _ := strconv.Atoi(string(s[i]))
		x = x*10 + tmp
	}
	ng := -1
	ok := 4000000000000000
	for abs(ng-ok) > 1 {
		k := (ng + ok) / 2
		if (a+k)*1000+b*2000+c*3000+d*4000 <= (n+k)*x {
			ok = k
		} else {
			ng = k
		}
	}
	fmt.Println(ok)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
