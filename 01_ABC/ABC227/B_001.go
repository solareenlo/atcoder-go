package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	for i := 0; i < n; i++ {
		var s int
		fmt.Scan(&s)
		ok := true
		for d := 1; d*7+3 <= s; d++ {
			if (s-3*d)%(4*d+3) == 0 {
				ok = false
			}
		}
		if ok {
			res++
		}
	}
	fmt.Println(res)
}
