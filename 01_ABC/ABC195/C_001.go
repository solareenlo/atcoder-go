package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	for i := 1000; i < n+1; i *= 1000 {
		res += n - i + 1
	}

	fmt.Println(res)
}
