package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	sum := 0
	if k-a >= 0 {
		sum += a
	} else {
		sum += k
	}

	k -= a
	k -= b

	if k > 0 {
		if k-c < 0 {
			sum -= k
		} else {
			sum -= c
		}
	}

	fmt.Println(sum)
}
