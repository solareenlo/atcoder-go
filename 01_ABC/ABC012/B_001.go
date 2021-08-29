package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%02d:%02d:%02d\n", n/3600, (n%3600)/60, n%60)
}
