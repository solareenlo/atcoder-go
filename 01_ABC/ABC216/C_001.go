package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := ""
	for n != 0 {
		if n%2 != 0 {
			res = "A" + res
		}
		n /= 2
		res = "B" + res
	}
	fmt.Println(res)
}
