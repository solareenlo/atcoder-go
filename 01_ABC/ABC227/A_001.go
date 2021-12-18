package main

import "fmt"

func main() {
	var n, k, a int
	fmt.Scan(&n, &k, &a)

	fmt.Println((a+k-2)%n + 1)
}
