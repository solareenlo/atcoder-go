package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		d := 0
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a%2 == 1 {
			d = 100
		}
		if b*10 >= d {
			d = 0
		} else {
			d -= b * 10
			b = 0
		}
		if b%2 == 1 {
			d += 10
		}
		if c >= d {
			d = 0
		}
		if c%2 == 1 {
			d = 1
		}
		if d == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
