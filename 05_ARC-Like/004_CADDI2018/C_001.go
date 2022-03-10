package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	ans := 1
	if n == 1 {
		ans = p
	} else {
		for i := 2; i*i <= p; i++ {
			var u int
			for u = 0; (p % i) == 0; u++ {
				p /= i
			}
			for u /= n; u > 0; u-- {
				ans *= i
			}
		}
	}
	fmt.Println(ans)
}
