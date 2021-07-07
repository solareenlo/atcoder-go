package main

import "fmt"

func main() {
	var n, l int64
	fmt.Scan(&n, &l)
	px := -l - 1
	p := 2 * -l
	for i := 0; i < int(n); i++ {
		var x int64
		fmt.Scan(&x)
		if x-px > l {
			p += 2 * l
			px = x
			if px-l > p {
				p = px - l
			}
		} else {
			if x-l > p {
				p = x - l
			}
		}
		if p >= px {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
