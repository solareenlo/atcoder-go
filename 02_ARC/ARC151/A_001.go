package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	var s, t string
	fmt.Scan(&n, &s, &t)
	d := 0
	for i := 0; i < n; i++ {
		A, _ := strconv.Atoi(string(s[i]))
		B, _ := strconv.Atoi(string(t[i]))
		d += A - B
	}
	if (d & 1) != 0 {
		fmt.Println(-1)
		return
	}
	u := make([]string, n)
	for n > 0 {
		n--
		if d < 0 && s[n] < t[n] {
			d += 2
			u[n] = "1"
		} else if d > 0 && s[n] > t[n] {
			d -= 2
			u[n] = "1"
		} else {
			u[n] = "0"
		}
	}
	fmt.Println(strings.Join(u, ""))
}
