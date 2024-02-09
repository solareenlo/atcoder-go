package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	k0 := 0
	k1 := 2000000000
	for k0+1 < k1 {
		k := (k0 + k1) / 2
		if k*(k+1)/2 <= n {
			k0 = k
		} else {
			k1 = k
		}
	}

	fmt.Println(k0)
}
