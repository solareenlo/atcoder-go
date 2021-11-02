package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	res := 0
	if a == 0 {
		res = 0
	} else if n <= a {
		res = n
	} else if n <= a+b {
		res = a
	} else if n%(a+b) == 0 {
		res = a * (n / (a + b))
	} else if n%(a+b) <= a {
		res = a*(n/(a+b)) + (n % (a + b))
	} else {
		res = a*(n/(a+b)) + a
	}

	fmt.Println(res)
}
