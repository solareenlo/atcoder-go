package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	if n%2 == 0 {
		tmp := 10
		for tmp <= n {
			res += n / tmp
			tmp *= 5
		}
	}

	fmt.Println(res)
}
