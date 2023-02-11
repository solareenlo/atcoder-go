package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for k := 1; k <= n; k++ {
		d := n / k
		e := d * (d + 1) / 2
		sum += k * e
	}
	fmt.Println(sum)
}
