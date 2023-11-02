package main

import "fmt"

var a [111]int
var i int

func f() {
	x := i
	if a[i] != 0 {
		i = a[i]
		f()
	}
	fmt.Printf("%d ", x)
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	for i = 0; i < m; i++ {
		var t int
		fmt.Scan(&t)
		a[t] = t + 1
	}
	for i = 1; i <= n; i++ {
		f()
	}
}
