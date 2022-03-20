package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(n/10000 + (n%10000)/5000 + (n%5000)/1000)
}
