package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var k int
	fmt.Scan(&k)
	k %= 118375
	for ; k > 0; k-- {
		a := n
		b := 0
		for a != 0 {
			b += a % 10
			a /= 10
		}
		n += b
		n %= 100000
	}
	fmt.Println(n)
}
