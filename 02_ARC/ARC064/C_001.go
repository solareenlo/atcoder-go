package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	k := 0
	ans := 0
	for i := 1; i < n+1; i++ {
		var a int
		fmt.Scan(&a)
		if k+a <= x {
			k = a
		} else {
			ans = ans + a + k - x
			k = x - k
		}
	}
	fmt.Println(ans)
}
