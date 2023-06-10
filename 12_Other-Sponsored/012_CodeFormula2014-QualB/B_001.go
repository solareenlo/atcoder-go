package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var a [2]int
	for i := 0; i < len(s); i++ {
		a[(len(s)-i)&1] += int(s[i] - '0')
	}
	fmt.Println(a[0], a[1])
}
