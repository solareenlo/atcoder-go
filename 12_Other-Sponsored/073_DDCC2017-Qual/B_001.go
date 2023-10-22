package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	sum := a*1728 + b*144 + c*12 + d
	fmt.Println(sum)
}
