package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := map[int]int{}
	sum := 0
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		sum += tmp
		a[tmp]++
	}

	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var b, c int
		fmt.Scan(&b, &c)
		if _, ok := a[b]; ok {
			sum += a[b] * (c - b)
			a[c] += a[b]
			a[b] = 0
		}
		fmt.Println(sum)
	}
}
