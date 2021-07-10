package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(fmt.Sprintf("%.6f", b/100*a))
}
