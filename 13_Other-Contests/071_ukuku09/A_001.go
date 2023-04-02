package main

import "fmt"

func main() {
	var a [200]int
	a[0], a[1] = 0, 1
	f := 1
	for a[f]+a[f-1] <= 1e18 {
		a[f+1] = a[f] + a[f-1]
		f++
	}
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		var x int
		fmt.Scan(&x)
		for i := f; i > 1; i-- {
			if x > a[i-1] && (x-a[i-1])%a[i] < 1 {
				fmt.Println(1, (x-a[i-1])/a[i])
				break
			}
		}
	}
}
