package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	fmt.Println((t/100 + t/10%10 + t%10) * 111)
}
