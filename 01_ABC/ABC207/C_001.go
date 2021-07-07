package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var l = make([]int, n)
	var r = make([]int, n)

	res := 0
	for i := 0; i < n; i++ {
		var t, l, r int
		fmt.Scan(&t, &l[i], &r[i])
		l[i] *= 2
		r[i] *= 2
		if t == 2 || t == 4 {
			r[i]--
		}
		if t == 3 || t == 4 {
			l[i]++
		}
		for j := 0; j < i; j++ {
			if l[i] <= r[j] && l[j] <= r[i] {
				res++
			}
		}
	}

	fmt.Println(res)
}
