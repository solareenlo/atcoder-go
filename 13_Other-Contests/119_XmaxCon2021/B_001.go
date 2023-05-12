package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(m+n+1, (m*n+m+n)/2+2)
}
