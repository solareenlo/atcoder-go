package main

import "fmt"

func main() {
	var sum, n, p, a int
	fmt.Scan(&n)
	for i := 1 << 17; i > 0; i >>= 1 {
		if n&i != 0 {
			fmt.Printf("? %d %d\n", p, p+i)
			fmt.Scan(&a)
			sum += a
			p += i
		}
	}
	fmt.Printf("! %d\n", sum)
}
