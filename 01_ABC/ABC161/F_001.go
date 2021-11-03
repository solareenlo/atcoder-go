package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 1
	for i := 2; i*i <= n; i++ {
		if (n-1)%i == 0 {
			res++
			if i*i != n-1 {
				res++
			}
		}
		h := n
		for h%i == 0 {
			h /= i
		}
		if h%i == 1 && n%i == 0 {
			res++
		}
	}

	if n > 2 {
		res++
	}
	fmt.Println(res)
}
