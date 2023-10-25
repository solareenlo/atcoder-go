package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(((N/2)*2)*(N/2-1) + N%2)
}
