package main

import "fmt"

func main() {
	var p int
	fmt.Scan(&p)

	res := 0
	i := 2

	for {
		res += p % i
		p /= i
		i++
		if p == 0 {
			break
		}
	}

	fmt.Println(res)
}
