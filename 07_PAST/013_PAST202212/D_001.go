package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var s string
	fmt.Scan(&s)

	a := make([]int, n)
	f := 0
	i := 0
	for _, c := range s {
		a[i]++
		if c == '+' {
			a[i] += f
			f = 0
		}
		if c == '-' {
			f += a[i]
			a[i] = 0
		}
		i = (i + 1) % n
	}

	for _, e := range a {
		fmt.Println(e)
	}
}
