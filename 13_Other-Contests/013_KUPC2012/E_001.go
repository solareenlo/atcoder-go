package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	c := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}

	rand.Seed(time.Now().UnixNano())
	const mod = 998244353
	e := rand.Intn(mod)%n + 1
	for i := 0; i < 1000; i++ {
		fmt.Println(e)
		var f int
		fmt.Scan(&f)
		if c[e-1][f-1] == 'x' {
			e = rand.Intn(mod)%n + 1
		}
	}
}
