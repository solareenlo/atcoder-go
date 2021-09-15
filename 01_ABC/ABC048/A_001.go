package main

import "fmt"

func main() {
	var a, s, c string
	fmt.Scan(&a, &s, &c)
	fmt.Println(string(a[0]) + string(s[0]) + string(c[0]))
}
