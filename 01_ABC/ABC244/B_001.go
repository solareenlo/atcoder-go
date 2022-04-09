package main

import "fmt"

func main() {
	var n int
	var t string
	fmt.Scan(&n, &t)

	cnt := 0
	x := [4]int{}
	for i := 0; i < n; i++ {
		if t[i] == 'S' {
			x[cnt]++
		}
		if t[i] == 'R' {
			cnt++
		}
		cnt %= 4
	}

	fmt.Println(x[0]-x[2], x[3]-x[1])
}
