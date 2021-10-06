package main

import "fmt"

func main() {
	var n, d, x, a int
	fmt.Scan(&n, &d, &x)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		cnt := 0
		for {
			if a*cnt+1 <= d {
				cnt++
			} else {
				break
			}
		}
		sum += cnt
	}
	fmt.Println(sum + x)
}
