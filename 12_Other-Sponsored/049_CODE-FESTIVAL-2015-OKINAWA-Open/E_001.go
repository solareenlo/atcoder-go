package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	a, c := W, 1
	for a%2 == 0 {
		a /= 2
	}
	for c < H {
		c *= 2
	}

	if a != 1 || H >= c-1 {
		d := 1
		for d < H*W {
			d *= 2
		}
		fmt.Println(d - 1)
	} else {
		fmt.Println((c - 1) * W)
	}
}
