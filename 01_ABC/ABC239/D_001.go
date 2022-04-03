package main

import "fmt"

func main() {
	prime := make([]int, 202)
	for i := 2; i <= 200; i++ {
		prime[i] = 1
	}
	for p := 2; p*p <= 200; p++ {
		if prime[p] != 0 {
			for i := p * p; i <= 200; i += p {
				prime[i] = 0
			}
		}
	}

	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	for x := a; x <= b; x++ {
		flag := 1
		for y := c; y <= d; y++ {
			if prime[x+y] == 0 {
				flag &= 1
			} else {
				flag &= 0
			}
		}
		if flag != 0 {
			fmt.Println("Takahashi")
			return
		}
	}
	fmt.Println("Aoki")
}
